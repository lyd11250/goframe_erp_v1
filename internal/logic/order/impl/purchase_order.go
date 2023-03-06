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

type sPurchaseOrder struct{}

var PurchaseOrder *sPurchaseOrder

func (s *sPurchaseOrder) InitOrderItem(ctx context.Context, in model.InitOrderItemInput) (err error) {
	orderInfoOutput, err := s.GetOrderInfo(ctx, model.GetOrderInfoInput{OrderNo: in.OrderNo})
	if err != nil {
		return err
	}
	var purchaseOrder entity.PurchaseOrder
	err = gconv.Struct(orderInfoOutput.Order, &purchaseOrder)
	if err != nil {
		return err
	}
	// 获取供应商信息
	supplierEnabled, err := service.Supplier().CheckSupplierEnabled(ctx, purchaseOrder.SupplierId)
	if err != nil {
		return err
	}
	if !supplierEnabled {
		return gerror.NewCode(gcode.CodeInvalidParameter, "供应商不可用")
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
		// 检查商品是否属于供应商
		count, err := dao.GoodsSupplierRel.Ctx(ctx).Count(g.Map{
			dao.GoodsSupplierRel.Columns().GoodsId:    orderItem.GoodsId,
			dao.GoodsSupplierRel.Columns().SupplierId: purchaseOrder.SupplierId,
		})
		if err != nil {
			return err
		}
		if count == 0 {
			return gerror.NewCodef(gcode.CodeInvalidParameter, "商品(%v)不属于供应商", orderItem.GoodsName)
		}

		in.Items[i].Status = consts.OrderStatusProcessing
		purchaseOrder.OrderQuantity += orderItem.Quantity
		purchaseOrder.OrderAmount += orderItem.Amount

		// 插入订单项
		_, err = dao.OrderItem.Ctx(ctx).OmitEmpty().Insert(in.Items[i])
		if err != nil {
			return err
		}
	}

	purchaseOrder.OrderStatus = consts.OrderStatusProcessing

	// 更新订单信息
	_, err = dao.PurchaseOrder.Ctx(ctx).OmitEmpty().Where(dao.PurchaseOrder.Columns().OrderNo, in.OrderNo).Update(purchaseOrder)
	return
}

func (s *sPurchaseOrder) CompleteOrder(ctx context.Context, in model.CompleteOrderInput) (err error) {
	// 获取订单信息
	orderInfoOutput, err := s.GetOrderInfo(ctx, model.GetOrderInfoInput{OrderNo: &in.OrderNo})
	if err != nil {
		return err
	}
	var purchaseOrder entity.PurchaseOrder
	var orderItems = orderInfoOutput.Items
	err = gconv.Struct(orderInfoOutput.Order, &purchaseOrder)
	if err != nil {
		return err
	}

	// 检查订单状态
	if purchaseOrder.OrderStatus != consts.OrderStatusProcessing {
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
	_, err = dao.PurchaseOrder.Ctx(ctx).
		Where(dao.PurchaseOrder.Columns().OrderNo, in.OrderNo).
		Data(dao.PurchaseOrder.Columns().Notes, in.Notes).
		Update()
	return
}

func (s *sPurchaseOrder) CompleteOrderItem(ctx context.Context, in model.CompleteOrderItemInput) (err error) {
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
	var purchaseOrder entity.PurchaseOrder
	var orderItems = orderInfoOutput.Items
	err = gconv.Struct(orderInfoOutput.Order, &purchaseOrder)
	if err != nil {
		return err
	}

	// 检查订单状态
	if purchaseOrder.OrderStatus != consts.OrderStatusProcessing {
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
		purchaseOrder.CompleteTime = currentTime
		purchaseOrder.CompleteUser = userInfo.UserId
		purchaseOrder.CompleteUserName = userInfo.UserRealName
		purchaseOrder.OrderStatus = consts.OrderStatusDone
		_, err = dao.PurchaseOrder.Ctx(ctx).
			Where(dao.PurchaseOrder.Columns().OrderNo, in.OrderNo).
			Update(purchaseOrder)
	}
	return
}

func (s *sPurchaseOrder) GetOrderInfo(ctx context.Context, in model.GetOrderInfoInput) (out model.GetOrderInfoOutput, err error) {
	// 获取订单信息
	result, err := dao.PurchaseOrder.Ctx(ctx).OmitNil().One(in)
	if err != nil {
		return model.GetOrderInfoOutput{}, err
	}
	// 订单不存在
	if result.IsEmpty() {
		return model.GetOrderInfoOutput{}, gerror.NewCode(gcode.CodeNotFound, "订单不存在")
	}
	// 订单存在，获取订单明细
	var order entity.PurchaseOrder
	err = result.Struct(&order)
	if err != nil {
		return model.GetOrderInfoOutput{}, err
	}
	out.Order = gconv.MapDeep(order)
	err = dao.OrderItem.Ctx(ctx).
		Where(dao.OrderItem.Columns().OrderNo, order.OrderNo).
		Scan(&out.Items)
	return
}

func (s *sPurchaseOrder) GetOrderList(ctx context.Context, in model.GetOrderListInput) (out model.GetOrderListOutput, err error) {
	var orderList []entity.PurchaseOrder
	err = dao.PurchaseOrder.Ctx(ctx).
		OrderDesc(dao.PurchaseOrder.Columns().CreateTime).
		Page(in.Page, in.PageSize).
		Scan(&orderList)
	if err != nil {
		return
	}
	err = gconv.Struct(orderList, &out.List)
	if err != nil {
		return
	}
	out.Total, err = dao.PurchaseOrder.Ctx(ctx).Count()
	if err != nil {
		return
	}
	out.Pages = out.Total / in.PageSize
	if out.Total%in.PageSize > 0 {
		out.Pages++
	}
	return
}

func (s *sPurchaseOrder) CreateOrder(ctx context.Context, in model.CreateOrderInput) (out model.CreateOrderOutput, err error) {
	currentTime := gtime.Now()
	// 获取供应商信息
	if in.PartyId == nil {
		return model.CreateOrderOutput{}, gerror.NewCode(gcode.CodeInvalidParameter, "客户ID不能为空")
	}
	supplierResult, err := service.Supplier().GetSupplierById(ctx, model.GetSupplierByIdInput{SupplierId: *in.PartyId})
	if err != nil {
		return model.CreateOrderOutput{}, err
	}
	if supplierResult.SupplierStatus != consts.StatusEnabled {
		return out, gerror.NewCode(gcode.CodeInvalidParameter, "供应商不可用")
	}
	// 初始化采购单信息
	userInfo, err := service.User().GetUserById(ctx, model.GetUserByIdInput{UserId: redis.Ctx(ctx).CheckLogin()})
	if err != nil {
		return model.CreateOrderOutput{}, err
	}
	order := entity.PurchaseOrder{
		OrderNo:        service.Order().GenerateOrderNo(*in.OrderType, currentTime),
		SupplierId:     supplierResult.SupplierId,
		SupplierName:   supplierResult.SupplierName,
		OrderAmount:    0,
		OrderQuantity:  0,
		CreateTime:     currentTime,
		CreateUser:     userInfo.UserId,
		CreateUserName: userInfo.UserRealName,
		OrderStatus:    consts.OrderStatusInit,
		Notes:          in.Notes,
	}
	order.OrderId, err = dao.PurchaseOrder.Ctx(ctx).OmitEmpty().InsertAndGetId(order)
	if err != nil {
		return
	}
	out.Order = gconv.MapDeep(order)
	return
}

func (s *sPurchaseOrder) CancelCreateOrder(ctx context.Context, in model.CancelCreateOrderInput) (err error) {
	// 获取订单信息
	orderResult, err := dao.PurchaseOrder.Ctx(ctx).OmitNil().One(in)
	if err != nil {
		return err
	}
	if orderResult.IsEmpty() {
		return gerror.NewCode(gcode.CodeNotFound, "订单不存在")
	}
	var order entity.PurchaseOrder
	err = orderResult.Struct(&order)
	if err != nil {
		return err
	}
	// 订单状态不是初始化状态
	if order.OrderStatus != consts.OrderStatusInit {
		return gerror.NewCode(gcode.CodeInvalidParameter, "订单状态错误，无法取消")
	}
	// 删除订单及订单项
	_, err = dao.PurchaseOrder.Ctx(ctx).Delete(dao.PurchaseOrder.Columns().OrderId, order.OrderId)
	if err != nil {
		return err
	}
	_, err = dao.OrderItem.Ctx(ctx).Delete(dao.OrderItem.Columns().OrderNo, order.OrderNo)
	return
}

func (s *sPurchaseOrder) CancelOrder(ctx context.Context, in model.CancelOrderInput) (err error) {
	// 获取订单信息
	orderInfoOutput, err := s.GetOrderInfo(ctx, model.GetOrderInfoInput{OrderNo: &in.OrderNo})
	if err != nil {
		return err
	}
	var order entity.PurchaseOrder
	err = gconv.Struct(orderInfoOutput.Order, &order)
	if err != nil {
		return err
	}
	// 订单状态不是处理中状态
	if order.OrderStatus != consts.OrderStatusProcessing {
		return gerror.NewCode(gcode.CodeInvalidParameter, "订单状态错误，无法取消")
	}

	// 更新订单状态
	_, err = dao.PurchaseOrder.Ctx(ctx).
		Where(dao.PurchaseOrder.Columns().OrderNo, in.OrderNo).
		Data(g.Map{
			dao.PurchaseOrder.Columns().OrderStatus: consts.OrderStatusCancel,
			dao.PurchaseOrder.Columns().Notes:       in.Notes,
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
