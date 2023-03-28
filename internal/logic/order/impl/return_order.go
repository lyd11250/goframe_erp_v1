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

type sReturnOrder struct{}

var ReturnOrder *sReturnOrder

func (s *sReturnOrder) InitOrderItem(ctx context.Context, in model.InitOrderItemInput) (err error) {
	// 获取订单信息
	orderInfoOutput, err := s.GetOrderInfo(ctx, model.GetOrderInfoInput{OrderNo: in.OrderNo})
	if err != nil {
		return err
	}
	var returnOrder entity.ReturnOrder
	err = gconv.Struct(orderInfoOutput.Order, &returnOrder)
	if err != nil {
		return err
	}

	// 获取源订单信息
	pOrderInfo, err := service.Order().Prefix(returnOrder.POrderNo[:4]).
		GetOrderInfo(ctx, model.GetOrderInfoInput{OrderNo: &returnOrder.POrderNo})
	if err != nil {
		return err
	}
	if pOrderInfo.Order["orderStatus"] != consts.OrderStatusDone {
		return gerror.NewCode(gcode.CodeInvalidParameter, "源订单未完成，无法退货")
	}

	// 商品退货数量
	goodsQuantityMap := make(map[int64]int)
	for _, orderItem := range in.Items {
		goodsQuantityMap[orderItem.GoodsId] += orderItem.Quantity
	}

	for goodsId, quantity := range goodsQuantityMap {
		// 检查商品
		enabled, err := service.Goods().CheckGoodsEnabled(ctx, model.CheckGoodsEnabledInput{GoodsId: goodsId})
		if err != nil {
			return err
		}
		if !enabled.Enabled {
			return gerror.NewCode(gcode.CodeInvalidParameter, "商品不可用")
		}

		// 检查商品采购/销售数量
		var pOrderQuantity int
		for _, pOrderItem := range pOrderInfo.Items {
			if pOrderItem.GoodsId == goodsId {
				pOrderQuantity += pOrderItem.Quantity
			}
		}
		if pOrderQuantity < quantity {
			return gerror.NewCode(gcode.CodeInvalidParameter, "商品退货数量大于源订单数量")
		}

		// 检查已退货数量
		var returnQuantity float64
		// 与源订单关联的、已完成的退货订单
		relReturnOrderNoResult, err := dao.ReturnOrder.Ctx(ctx).
			Fields(dao.ReturnOrder.Columns().OrderNo).
			Where(dao.ReturnOrder.Columns().POrderNo, returnOrder.POrderNo).
			Where(dao.ReturnOrder.Columns().OrderStatus, consts.OrderStatusDone).
			All()
		if err != nil {
			return err
		}
		returnQuantity, err = dao.OrderItem.Ctx(ctx).
			WhereIn(dao.OrderItem.Columns().OrderNo, relReturnOrderNoResult.Array("order_no")).
			Where(dao.OrderItem.Columns().GoodsId, goodsId).
			Sum(dao.OrderItem.Columns().Quantity)
		if err != nil {
			return err
		}
		if pOrderQuantity < quantity+int(returnQuantity) {
			return gerror.NewCode(gcode.CodeInvalidParameter, "商品退货数量大于源订单数量")
		}
	}

	for i := range in.Items {
		in.Items[i].Status = consts.OrderStatusProcessing
		returnOrder.OrderQuantity += in.Items[i].Quantity
		returnOrder.OrderAmount += in.Items[i].Amount

		// 插入订单项
		_, err = dao.OrderItem.Ctx(ctx).OmitEmpty().Insert(in.Items[i])
		if err != nil {
			return err
		}
	}

	returnOrder.OrderStatus = consts.OrderStatusProcessing

	// 更新订单信息
	_, err = dao.ReturnOrder.Ctx(ctx).OmitEmpty().Where(dao.PurchaseOrder.Columns().OrderNo, in.OrderNo).Update(returnOrder)
	return
}

func (s *sReturnOrder) CompleteOrder(ctx context.Context, in model.CompleteOrderInput) (err error) {
	// 获取订单信息
	orderInfoOutput, err := s.GetOrderInfo(ctx, model.GetOrderInfoInput{OrderNo: &in.OrderNo})
	if err != nil {
		return err
	}
	var returnOrder entity.ReturnOrder
	var orderItems = orderInfoOutput.Items
	err = gconv.Struct(orderInfoOutput.Order, &returnOrder)
	if err != nil {
		return err
	}

	// 检查订单状态
	if returnOrder.OrderStatus != consts.OrderStatusProcessing {
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

func (s *sReturnOrder) CompleteOrderItem(ctx context.Context, in model.CompleteOrderItemInput) (err error) {
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
	var returnOrder entity.ReturnOrder
	var orderItems = orderInfoOutput.Items
	err = gconv.Struct(orderInfoOutput.Order, &returnOrder)
	if err != nil {
		return err
	}

	// 检查订单状态
	if returnOrder.OrderStatus != consts.OrderStatusProcessing {
		return gerror.NewCode(gcode.CodeInvalidParameter, "订单状态不正确")
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
		returnOrder.CompleteTime = currentTime
		returnOrder.CompleteUser = userInfo.UserId
		returnOrder.CompleteUserName = userInfo.UserRealName
		returnOrder.OrderStatus = consts.OrderStatusDone
		_, err = dao.ReturnOrder.Ctx(ctx).
			Where(dao.ReturnOrder.Columns().OrderNo, in.OrderNo).
			Update(returnOrder)
	}
	return
}

func (s *sReturnOrder) GetOrderInfo(ctx context.Context, in model.GetOrderInfoInput) (out model.GetOrderInfoOutput, err error) {
	// 获取订单信息
	result, err := dao.ReturnOrder.Ctx(ctx).OmitNil().One(in)
	if err != nil {
		return model.GetOrderInfoOutput{}, err
	}

	// 订单不存在
	if result.IsEmpty() {
		return model.GetOrderInfoOutput{}, gerror.NewCode(gcode.CodeNotFound, "订单不存在")
	}

	// 订单存在，获取订单明细
	var order entity.ReturnOrder
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

	// 已完成退货出/入库数量

	return
}

func (s *sReturnOrder) GetOrderList(ctx context.Context, in model.GetOrderListInput) (out model.GetOrderListOutput, err error) {
	var orderList []entity.ReturnOrder
	err = dao.ReturnOrder.Ctx(ctx).
		OmitEmpty().
		Where(dao.ReturnOrder.Columns().OrderStatus, in.OrderStatus).
		Where(dao.ReturnOrder.Columns().OrderType, in.OrderType).
		OrderDesc(dao.ReturnOrder.Columns().CreateTime).
		Page(in.Page, in.PageSize).
		Scan(&orderList)
	if err != nil {
		return
	}
	err = gconv.Struct(orderList, &out.List)
	if err != nil {
		return
	}
	out.Total, err = dao.ReturnOrder.Ctx(ctx).
		OmitEmpty().
		Where(dao.ReturnOrder.Columns().OrderStatus, in.OrderStatus).
		Where(dao.ReturnOrder.Columns().OrderType, in.OrderType).
		Count()
	if err != nil {
		return
	}
	out.Pages = out.Total / in.PageSize
	if out.Total%in.PageSize > 0 {
		out.Pages++
	}
	return
}

func (s *sReturnOrder) CreateOrder(ctx context.Context, in model.CreateOrderInput) (out model.CreateOrderOutput, err error) {
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

	// 若源订单未完成，则不允许创建退货单
	if pOrderInfo.Order["orderStatus"] != consts.OrderStatusDone {
		return model.CreateOrderOutput{}, gerror.NewCode(gcode.CodeInvalidParameter, "源订单未完成，不允许创建退货单")
	}

	// 初始化退货单信息
	userInfo, err := service.User().GetUserById(ctx, model.GetUserByIdInput{UserId: redis.Ctx(ctx).CheckLogin()})
	if err != nil {
		return model.CreateOrderOutput{}, err
	}
	order := entity.ReturnOrder{
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
	order.OrderId, err = dao.ReturnOrder.Ctx(ctx).InsertAndGetId(order)
	if err != nil {
		return
	}
	out.Order = gconv.MapDeep(order)
	return
}

func (s *sReturnOrder) CancelCreateOrder(ctx context.Context, in model.CancelCreateOrderInput) (err error) {
	// 获取订单信息
	orderResult, err := dao.ReturnOrder.Ctx(ctx).One(in)
	if err != nil {
		return err
	}
	if orderResult.IsEmpty() {
		return gerror.NewCode(gcode.CodeNotFound, "订单不存在")
	}
	var order entity.ReturnOrder
	err = orderResult.Struct(&order)
	if err != nil {
		return err
	}
	// 订单状态不是初始化状态
	if order.OrderStatus != consts.OrderStatusInit {
		return gerror.NewCode(gcode.CodeInvalidParameter, "订单状态错误，无法取消")
	}
	// 删除订单及订单项
	_, err = dao.ReturnOrder.Ctx(ctx).Delete(dao.ReturnOrder.Columns().OrderId, order.OrderId)
	if err != nil {
		return err
	}
	_, err = dao.OrderItem.Ctx(ctx).Delete(dao.OrderItem.Columns().OrderNo, order.OrderNo)
	return
}

func (s *sReturnOrder) CancelOrder(ctx context.Context, in model.CancelOrderInput) (err error) {
	// 获取订单信息
	orderInfoOutput, err := s.GetOrderInfo(ctx, model.GetOrderInfoInput{OrderNo: &in.OrderNo})
	if err != nil {
		return err
	}
	var order entity.ReturnOrder
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
	_, err = dao.ReturnOrder.Ctx(ctx).
		Where(dao.ReturnOrder.Columns().OrderNo, in.OrderNo).
		Data(g.Map{
			dao.ReturnOrder.Columns().OrderStatus: consts.OrderStatusCancel,
			dao.ReturnOrder.Columns().Notes:       in.Notes,
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
