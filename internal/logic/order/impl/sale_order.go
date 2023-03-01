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

type sSaleOrder struct{}

var SaleOrder *sSaleOrder

func (s *sSaleOrder) GetOrderInfo(ctx context.Context, in model.GetOrderInfoInput) (out model.GetOrderInfoOutput, err error) {
	// 获取订单信息
	result, err := dao.SaleOrder.Ctx(ctx).OmitNil().One(in)
	if err != nil {
		return model.GetOrderInfoOutput{}, err
	}
	// 订单不存在
	if result.IsEmpty() {
		return model.GetOrderInfoOutput{}, gerror.NewCode(gcode.CodeNotFound, "订单不存在")
	}
	// 订单存在，获取订单明细
	out.Order = result.Map()
	err = dao.OrderItem.Ctx(ctx).
		Where(dao.OrderItem.Columns().OrderNo, out.Order[dao.SaleOrder.Columns().OrderNo]).
		Scan(&out.Items)
	return
}

func (s *sSaleOrder) GetOrderList(ctx context.Context, in model.GetOrderListInput) (out model.GetOrderListOutput, err error) {
	result, err := dao.SaleOrder.Ctx(ctx).
		Page(in.Page, in.PageSize).
		All()
	if err != nil {
		return
	}
	out.List = result.List()
	out.Total, err = dao.SaleOrder.Ctx(ctx).Count()
	if err != nil {
		return
	}
	out.Pages = out.Total / in.PageSize
	if out.Total%in.PageSize > 0 {
		out.Pages++
	}
	return
}

func (s *sSaleOrder) CreateOrder(ctx context.Context, in model.CreateOrderInput) (out model.CreateOrderOutput, err error) {
	currentTime := gtime.Now()
	// 获取客户信息
	if in.PartyId == nil {
		return model.CreateOrderOutput{}, gerror.NewCode(gcode.CodeInvalidParameter, "客户ID不能为空")
	}
	customerResult, err := service.Customer().GetCustomerById(ctx, model.GetCustomerByIdInput{CustomerId: *in.PartyId})
	if err != nil {
		return model.CreateOrderOutput{}, err
	}
	// 初始化销售单信息
	order := entity.SaleOrder{
		OrderNo:       service.Order().GenerateOrderNo(*in.OrderType, currentTime),
		CustomerId:    customerResult.CustomerId,
		CustomerName:  customerResult.CustomerName,
		OrderAmount:   0,
		OrderQuantity: 0,
		CreateTime:    currentTime,
		CreateUser:    redis.Ctx(ctx).CheckLogin(),
		OrderStatus:   consts.OrderStatusInit,
	}
	order.OrderId, err = dao.SaleOrder.Ctx(ctx).InsertAndGetId(order)
	if err != nil {
		return
	}
	out.Order = gconv.MapDeep(order)
	return
}

func (s *sSaleOrder) CancelCreateOrder(ctx context.Context, in model.CancelCreateOrderInput) (err error) {
	// 获取订单信息
	orderResult, err := dao.SaleOrder.Ctx(ctx).OmitNil().One(in)
	if err != nil {
		return err
	}
	if orderResult.IsEmpty() {
		return gerror.NewCode(gcode.CodeNotFound, "订单不存在")
	}
	var order entity.SaleOrder
	err = orderResult.Struct(&order)
	if err != nil {
		return err
	}
	// 订单状态不是初始化状态
	if order.OrderStatus != consts.OrderStatusInit {
		return gerror.NewCode(gcode.CodeInvalidParameter, "订单状态错误，无法取消")
	}
	// 删除订单及订单项
	_, err = dao.SaleOrder.Ctx(ctx).Delete(dao.SaleOrder.Columns().OrderId, order.OrderId)
	if err != nil {
		return err
	}
	_, err = dao.OrderItem.Ctx(ctx).Delete(dao.OrderItem.Columns().OrderNo, order.OrderNo)
	return
}
