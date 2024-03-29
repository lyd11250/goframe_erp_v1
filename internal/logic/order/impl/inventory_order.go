package impl

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-erp-v1/internal/consts"
	"goframe-erp-v1/internal/dao"
	"goframe-erp-v1/internal/model"
	"goframe-erp-v1/internal/model/entity"
	"goframe-erp-v1/internal/service"
	"goframe-erp-v1/utility/redis"
)

type sInventoryOrder struct{}

var InventoryOrder *sInventoryOrder

func (s *sInventoryOrder) InitOrderItem(ctx context.Context, in model.InitOrderItemInput) (err error) {
	orderInfoOutput, err := s.GetOrderInfo(ctx, model.GetOrderInfoInput{OrderNo: in.OrderNo})
	if err != nil {
		return err
	}
	var inventoryOrder entity.InventoryOrder
	err = gconv.Struct(orderInfoOutput.Order, &inventoryOrder)
	if err != nil {
		return err
	}

	goodsQuantityMap := make(map[int64]int)
	for _, orderItem := range in.Items {
		goodsQuantityMap[orderItem.GoodsId] += orderItem.Quantity
	}

	for i, orderItem := range in.Items {
		// 检查商品
		checkGoodsEnabledOutput, err := service.Goods().CheckGoodsEnabled(ctx, model.CheckGoodsEnabledInput{GoodsId: orderItem.GoodsId})
		if err != nil {
			return err
		}
		if !checkGoodsEnabledOutput.Enabled {
			return gerror.NewCodef(gcode.CodeInvalidParameter, "商品(%v)不可用", orderItem.GoodsName)
		}

		// 源订单号
		pOrderNo := orderInfoOutput.Order["pOrderNo"]

		// 与源订单关联的已完成的出/入库单单号
		result, err := dao.InventoryOrder.Ctx(ctx).
			Fields(dao.InventoryOrder.Columns().OrderNo).
			All(g.Map{
				dao.InventoryOrder.Columns().POrderNo:    pOrderNo,
				dao.InventoryOrder.Columns().OrderStatus: consts.OrderStatusDone,
			})
		if err != nil {
			return err
		}
		var relInventoryOrderNoList = result.Array()

		// 已完成的出/入库单数量
		doneSum, err := dao.OrderItem.Ctx(ctx).
			WhereIn(dao.OrderItem.Columns().OrderNo, relInventoryOrderNoList).
			Where(dao.OrderItem.Columns().Status, consts.OrderStatusDone).
			Sum(dao.OrderItem.Columns().Quantity)
		if err != nil {
			return err
		}

		// 源订单采购/销售数量
		pOrderSum, err := dao.OrderItem.Ctx(ctx).
			Where(g.Map{
				dao.OrderItem.Columns().OrderNo: pOrderNo,
				dao.OrderItem.Columns().GoodsId: orderItem.GoodsId,
			}).Sum(dao.OrderItem.Columns().Quantity)
		if err != nil {
			return err
		}

		if pOrderSum == 0 {
			return gerror.NewCodef(gcode.CodeInvalidParameter, "商品(%v)未在源订单中", orderItem.GoodsName)
		}

		if pOrderSum-doneSum < float64(goodsQuantityMap[orderItem.GoodsId]) {
			return gerror.NewCodef(gcode.CodeInvalidParameter, "商品(%v)出/入库超量", orderItem.GoodsName)
		}

		in.Items[i].Status = consts.OrderStatusProcessing
		inventoryOrder.OrderQuantity += orderItem.Quantity
		inventoryOrder.OrderAmount += orderItem.Amount
	}

	inventoryOrder.OrderStatus = consts.OrderStatusProcessing

	// 插入订单项
	_, err = dao.OrderItem.Ctx(ctx).Insert(in.Items)
	// 更新订单信息
	_, err = dao.InventoryOrder.Ctx(ctx).OmitEmpty().Where(dao.InventoryOrder.Columns().OrderNo, in.OrderNo).Update(inventoryOrder)
	return
}

func (s *sInventoryOrder) CompleteOrder(ctx context.Context, in model.CompleteOrderInput) (err error) {
	// 获取订单信息
	orderInfoOutput, err := s.GetOrderInfo(ctx, model.GetOrderInfoInput{OrderNo: &in.OrderNo})
	if err != nil {
		return err
	}
	var inventoryOrder entity.InventoryOrder
	var orderItems = orderInfoOutput.Items
	err = gconv.Struct(orderInfoOutput.Order, &inventoryOrder)
	if err != nil {
		return err
	}

	// 检查订单状态
	if inventoryOrder.OrderStatus != consts.OrderStatusProcessing {
		return gerror.NewCode(gcode.CodeInvalidParameter, "订单状态不正确")
	}

	// 完成所有订单项
	for _, item := range orderItems {
		if item.Status == consts.OrderStatusProcessing {
			err = s.CompleteOrderItem(ctx, model.CompleteOrderItemInput{
				OrderNo:     in.OrderNo,
				OrderItemId: item.OrderItemId,
				Notes:       item.Notes,
			})
			if err != nil {
				return err
			}
		}
	}

	// 更新订单备注
	_, err = dao.InventoryOrder.Ctx(ctx).
		Where(dao.InventoryOrder.Columns().OrderNo, in.OrderNo).
		Data(dao.InventoryOrder.Columns().Notes, in.Notes).
		Update()
	return
}

func (s *sInventoryOrder) CompleteOrderItem(ctx context.Context, in model.CompleteOrderItemInput) (err error) {
	// 获取用户信息
	userInfo, err := service.User().GetUserById(ctx, model.GetUserByIdInput{UserId: redis.Ctx(ctx).CheckLogin()})
	if err != nil {
		return err
	}
	// 获取订单信息
	orderInfoOutput, err := s.GetOrderInfo(ctx, model.GetOrderInfoInput{OrderNo: &in.OrderNo})
	if err != nil {
		return err
	}
	var inventoryOrder entity.InventoryOrder
	var orderItems = orderInfoOutput.Items
	err = gconv.Struct(orderInfoOutput.Order, &inventoryOrder)
	if err != nil {
		return err
	}

	// 检查订单状态
	if inventoryOrder.OrderStatus != consts.OrderStatusProcessing {
		return gerror.NewCode(gcode.CodeInvalidParameter, "订单状态不正确")
	}

	// 源订单信息
	prefix := inventoryOrder.POrderNo[:4]
	pOrderInfo, err := service.Order().Prefix(prefix).GetOrderInfo(ctx, model.GetOrderInfoInput{OrderNo: &inventoryOrder.POrderNo})
	if err != nil {
		return err
	}
	// 获取源订单关联的已完成的出入库订单号
	result, err := dao.InventoryOrder.Ctx(ctx).
		Where(dao.InventoryOrder.Columns().POrderNo, inventoryOrder.POrderNo).
		Where(dao.InventoryOrder.Columns().OrderStatus, consts.OrderStatusDone).
		Fields(dao.InventoryOrder.Columns().OrderNo).
		All()
	if err != nil {
		return err
	}
	relOrderNoArray := result.Array()
	// 获取已完成订单项的出入库总数
	doneSum, err := dao.OrderItem.Ctx(ctx).
		WhereIn(dao.OrderItem.Columns().OrderNo, relOrderNoArray).
		Where(dao.OrderItem.Columns().Status, consts.OrderStatusDone).
		Sum(dao.OrderItem.Columns().Quantity)
	if err != nil {
		return err
	}

	// 检查订单项是否存在
	var orderItem entity.OrderItem
	var completedItemNum = 0
	for _, item := range orderItems {
		if item.OrderItemId == in.OrderItemId {
			orderItem = item
			break
		}
		if item.Status == consts.OrderStatusDone {
			completedItemNum++
		}
	}
	if orderItem.OrderItemId == 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "订单项不存在")
	}

	// 检查订单项状态
	if orderItem.Status != consts.OrderStatusProcessing {
		return gerror.NewCode(gcode.CodeInvalidParameter, "订单项状态不正确")
	}

	// 检查出入库数量
	for _, pOrderItem := range pOrderInfo.Items {
		if pOrderItem.GoodsId == orderItem.GoodsId {
			if orderItem.Quantity+int(doneSum) > pOrderItem.Quantity {
				return gerror.NewCodef(gcode.CodeInvalidParameter, "商品(%v)出入库超量", orderItem.GoodsName)
			}
		}
	}

	// 商品出入库
	var inventory entity.Inventory
	err = gconv.Struct(orderItem, &inventory)
	if err != nil {
		return err
	}
	switch inventoryOrder.OrderType {
	case consts.OrderTypeCGRK:
		_, err = service.Inventory().AddInventory(ctx, model.AddInventoryInput{
			Inventory: inventory,
		})
	case consts.OrderTypeXSCK:
		_, err = service.Inventory().ReduceInventory(ctx, model.ReduceInventoryInput{
			GoodsId:  orderItem.GoodsId,
			Quantity: orderItem.Quantity,
		})
	}
	if err != nil {
		return err
	}

	// 完成订单项
	currentTime := gtime.Now()
	orderItem.Status = consts.OrderStatusDone
	orderItem.CompleteTime = currentTime
	orderItem.CompleteUser = userInfo.UserId
	orderItem.CompleteUserName = userInfo.UserRealName
	orderItem.Notes = in.Notes
	_, err = dao.OrderItem.Ctx(ctx).
		Where(dao.OrderItem.Columns().OrderItemId, in.OrderItemId).
		Update(orderItem)
	if err != nil {
		return err
	}

	// 若订单的所有订单项全部完成，更新订单信息
	if completedItemNum == len(orderItems)-1 {
		inventoryOrder.CompleteTime = currentTime
		inventoryOrder.CompleteUser = userInfo.UserId
		inventoryOrder.CompleteUserName = userInfo.UserRealName
		inventoryOrder.OrderStatus = consts.OrderStatusDone
		_, err = dao.InventoryOrder.Ctx(ctx).
			Where(dao.InventoryOrder.Columns().OrderNo, in.OrderNo).
			Update(inventoryOrder)
	}
	return
}

func (s *sInventoryOrder) GetOrderInfo(ctx context.Context, in model.GetOrderInfoInput) (out model.GetOrderInfoOutput, err error) {
	// 获取订单信息
	result, err := dao.InventoryOrder.Ctx(ctx).OmitNil().One(in)
	if err != nil {
		return model.GetOrderInfoOutput{}, err
	}

	// 订单不存在
	if result.IsEmpty() {
		return model.GetOrderInfoOutput{}, gerror.NewCode(gcode.CodeNotFound, "订单不存在")
	}

	// 订单存在，获取订单明细
	var order entity.InventoryOrder
	err = result.Struct(&order)
	if err != nil {
		return model.GetOrderInfoOutput{}, err
	}
	out.Order = gconv.MapDeep(order)
	err = dao.OrderItem.Ctx(ctx).
		Where(dao.OrderItem.Columns().OrderNo, order.OrderNo).
		Scan(&out.Items)
	if err != nil {
		return model.GetOrderInfoOutput{}, err
	}

	return
}

func (s *sInventoryOrder) GetOrderList(ctx context.Context, in model.GetOrderListInput) (out model.GetOrderListOutput, err error) {
	var orderList []entity.InventoryOrder
	err = dao.InventoryOrder.Ctx(ctx).
		OmitEmpty().
		Where(dao.InventoryOrder.Columns().OrderStatus, in.OrderStatus).
		Where(dao.InventoryOrder.Columns().OrderType, in.OrderType).
		OrderDesc(dao.InventoryOrder.Columns().CreateTime).
		Page(in.Page, in.PageSize).
		Scan(&orderList)
	if err != nil {
		return
	}
	err = gconv.Struct(orderList, &out.List)
	if err != nil {
		return
	}
	out.Total, err = dao.InventoryOrder.Ctx(ctx).
		OmitEmpty().
		Where(dao.InventoryOrder.Columns().OrderStatus, in.OrderStatus).
		Count(dao.InventoryOrder.Columns().OrderType, in.OrderType)
	if err != nil {
		return
	}
	out.Pages = out.Total / in.PageSize
	if out.Total%in.PageSize > 0 {
		out.Pages++
	}
	return
}

func (s *sInventoryOrder) CreateOrder(ctx context.Context, in model.CreateOrderInput) (out model.CreateOrderOutput, err error) {
	currentTime := gtime.Now()
	// 获取源单据信息
	if in.POrderNo == nil {
		return model.CreateOrderOutput{}, gerror.NewCode(gcode.CodeInvalidParameter, "源单号不能为空")
	}
	prefix := *in.POrderNo
	prefix = prefix[:4]
	pOrderInfo, err := service.Order().Prefix(prefix).GetOrderInfo(ctx, model.GetOrderInfoInput{OrderNo: in.POrderNo})
	if err != nil {
		return model.CreateOrderOutput{}, err
	}
	out.POrder = gconv.MapDeep(pOrderInfo)

	// 若源订单未完成，则不允许创建出/入库单
	if pOrderInfo.Order["orderStatus"] != consts.OrderStatusDone {
		return model.CreateOrderOutput{}, gerror.NewCode(gcode.CodeInvalidParameter, "源订单未完成，不允许创建出/入库单")
	}

	// 初始化出/入库单信息
	userInfo, err := service.User().GetUserById(ctx, model.GetUserByIdInput{UserId: redis.Ctx(ctx).CheckLogin()})
	if err != nil {
		return model.CreateOrderOutput{}, err
	}
	order := entity.InventoryOrder{
		OrderNo:        service.Order().GenerateOrderNo(*in.OrderType, currentTime),
		OrderType:      *in.OrderType,
		POrderNo:       *in.POrderNo,
		OrderAmount:    0,
		OrderQuantity:  0,
		CreateTime:     currentTime,
		CreateUser:     userInfo.UserId,
		CreateUserName: userInfo.UserRealName,
		OrderStatus:    consts.OrderStatusInit,
		Notes:          in.Notes,
	}
	order.OrderId, err = dao.InventoryOrder.Ctx(ctx).InsertAndGetId(order)
	if err != nil {
		return
	}
	out.Order = gconv.MapDeep(order)
	return
}

func (s *sInventoryOrder) CancelCreateOrder(ctx context.Context, in model.CancelCreateOrderInput) (err error) {
	// 获取订单信息
	orderResult, err := dao.InventoryOrder.Ctx(ctx).OmitNil().One(in)
	if err != nil {
		return err
	}
	if orderResult.IsEmpty() {
		return gerror.NewCode(gcode.CodeNotFound, "订单不存在")
	}
	var order entity.InventoryOrder
	err = orderResult.Struct(&order)
	if err != nil {
		return err
	}
	// 订单状态不是初始化状态
	if order.OrderStatus != consts.OrderStatusInit {
		return gerror.NewCode(gcode.CodeInvalidParameter, "订单状态错误，无法取消")
	}
	// 删除订单及订单项
	_, err = dao.InventoryOrder.Ctx(ctx).Delete(dao.InventoryOrder.Columns().OrderId, order.OrderId)
	if err != nil {
		return err
	}
	_, err = dao.OrderItem.Ctx(ctx).Delete(dao.OrderItem.Columns().OrderNo, order.OrderNo)
	return
}

func (s *sInventoryOrder) CancelOrder(ctx context.Context, in model.CancelOrderInput) (err error) {
	// 获取订单信息
	orderInfoOutput, err := s.GetOrderInfo(ctx, model.GetOrderInfoInput{OrderNo: &in.OrderNo})
	if err != nil {
		return err
	}
	var order entity.InventoryOrder
	var items = orderInfoOutput.Items
	err = gconv.Struct(orderInfoOutput.Order, &order)
	if err != nil {
		return err
	}
	// 订单状态不是处理中状态
	if order.OrderStatus != consts.OrderStatusProcessing {
		return gerror.NewCode(gcode.CodeInvalidParameter, "订单状态错误，无法取消")
	}
	// 检查订单项状态
	for _, item := range items {
		if item.Status == consts.OrderStatusDone {
			return gerror.NewCode(gcode.CodeInvalidParameter, "订单项已完成，无法取消")
		}
	}
	// 更新订单状态
	_, err = dao.InventoryOrder.Ctx(ctx).
		Where(dao.InventoryOrder.Columns().OrderNo, in.OrderNo).
		Data(g.Map{
			dao.InventoryOrder.Columns().OrderStatus: consts.OrderStatusCancel,
			dao.InventoryOrder.Columns().Notes:       in.Notes,
		}).
		Update()
	if err != nil {
		return err
	}

	// 更新订单项状态
	_, err = dao.OrderItem.Ctx(ctx).
		Where(dao.OrderItem.Columns().OrderNo, in.OrderNo).
		Data(dao.OrderItem.Columns().Status, consts.OrderStatusCancel).
		Update()
	return
}
