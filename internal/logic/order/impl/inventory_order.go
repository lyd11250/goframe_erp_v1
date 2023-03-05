package impl

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
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
	// 获取源订单信息
	prefix := *in.OrderNo
	prefix = prefix[:4]
	getOrderInfoOutput, err := service.Order().Prefix(prefix).GetOrderInfo(ctx, model.GetOrderInfoInput{OrderNo: in.OrderNo})
	if err != nil {
		return err
	}

	var pOrderGoodsQuantityMap map[int64]*int
	for _, pOrderItem := range getOrderInfoOutput.Items {
		pOrderGoodsQuantityMap[pOrderItem.GoodsId] = &pOrderItem.Quantity
	}

	for _, orderItem := range in.Items {
		// 检查商品
		checkGoodsEnabledOutput, err := service.Goods().CheckGoodsEnabled(ctx, model.CheckGoodsEnabledInput{GoodsId: orderItem.GoodsId})
		if err != nil {
			return err
		}
		if !checkGoodsEnabledOutput.Enabled {
			return gerror.NewCodef(gcode.CodeInvalidParameter, "商品(%v)不可用", orderItem.GoodsName)
		}

		// 检查商品是否属于源订单，且数量正确
		quantity := pOrderGoodsQuantityMap[orderItem.GoodsId]
		if quantity == nil {
			return gerror.NewCodef(gcode.CodeInvalidParameter, "商品(%v)不属于源订单", orderItem.GoodsName)
		}
		if *quantity < orderItem.Quantity {
			return gerror.NewCodef(gcode.CodeInvalidParameter, "商品(%v)出入库数量大于源订单数量", orderItem.GoodsName)
		}

		inventoryOrder.OrderQuantity += orderItem.Quantity
		inventoryOrder.OrderAmount += orderItem.Amount
	}

	inventoryOrder.OrderStatus = consts.OrderStatusProcessing

	// 插入订单项
	_, err = dao.OrderItem.Ctx(ctx).Insert(in.Items)
	// 更新订单信息
	_, err = dao.PurchaseOrder.Ctx(ctx).OmitEmpty().Where(dao.PurchaseOrder.Columns().OrderNo, in.OrderNo).Update(inventoryOrder)
	return
}

var InventoryOrder *sInventoryOrder

func (s *sInventoryOrder) CompleteOrder(ctx context.Context, in model.CompleteOrderInput) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *sInventoryOrder) CompleteOrderItem(ctx context.Context, in model.CompleteOrderItemInput) (err error) {
	//TODO implement me
	panic("implement me")
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
		Where(dao.OrderItem.Columns().OrderNo, out.Order[dao.InventoryOrder.Columns().OrderNo]).
		Scan(&out.Items)
	return
}

func (s *sInventoryOrder) GetOrderList(ctx context.Context, in model.GetOrderListInput) (out model.GetOrderListOutput, err error) {
	var orderList []entity.InventoryOrder
	err = dao.InventoryOrder.Ctx(ctx).
		Page(in.Page, in.PageSize).
		Scan(&orderList)
	if err != nil {
		return
	}
	err = gconv.Struct(orderList, &out.List)
	if err != nil {
		return
	}
	out.Total, err = dao.InventoryOrder.Ctx(ctx).Count(dao.InventoryOrder.Columns().OrderType, in.OrderType)
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
	var pOrderType int
	// 销售出库单，获取销售单信息
	if *in.OrderType == consts.OrderTypeXSCK {
		pOrderType = consts.OrderTypeXSDD
	}
	// 采购入库单，获取采购单信息
	if *in.OrderType == consts.OrderTypeCGRK {
		pOrderType = consts.OrderTypeCGDD
	}
	pOrderInfo, err := service.Order().Type(pOrderType).GetOrderInfo(ctx, model.GetOrderInfoInput{OrderNo: in.POrderNo})
	if err != nil {
		return model.CreateOrderOutput{}, err
	}
	out.POrder = gconv.MapDeep(pOrderInfo)

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
	return
}
