package order

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"goframe-erp-v1/internal/consts"
	"goframe-erp-v1/internal/dao"
	"goframe-erp-v1/internal/model/entity"
	"goframe-erp-v1/internal/model/pojo"
	"goframe-erp-v1/internal/model/pojo/order"
	"goframe-erp-v1/internal/service"
	"goframe-erp-v1/utility/redis"
)

func (s *sOrder) GetSaleOrder(ctx context.Context, in order.GetSaleOrderInput) (out order.GetSaleOrderOutput, err error) {
	orderResult, err := dao.SaleOrder.Ctx(ctx).OmitNil().One(in)

	if err != nil {
		return order.GetSaleOrderOutput{}, err
	}
	if orderResult.IsEmpty() {
		return order.GetSaleOrderOutput{}, gerror.NewCode(gcode.CodeNotFound, "订单不存在")
	}
	err = orderResult.Struct(&out.Order)
	if err != nil {
		return order.GetSaleOrderOutput{}, err
	}
	err = dao.OrderItem.Ctx(ctx).Where(dao.OrderItem.Columns().OrderId, out.Order.OrderId).Scan(&out.Items)
	return
}

func (s *sOrder) GetSaleOrderList(ctx context.Context, in order.GetSaleOrderListInput) (out order.GetSaleOrderListOutput, err error) {
	err = dao.SaleOrder.Ctx(ctx).Page(in.Page, in.PageSize).Scan(&out.List)
	if err != nil {
		return order.GetSaleOrderListOutput{}, err
	}
	out.Total, err = dao.SaleOrder.Ctx(ctx).Count()
	if err != nil {
		return order.GetSaleOrderListOutput{}, err
	}
	out.Pages = out.Total / in.PageSize
	if out.Total%in.PageSize != 0 {
		out.Pages++
	}
	return
}

func (s *sOrder) CreateSaleOrder(ctx context.Context, in order.CreateSaleOrderInput) (out order.CreateSaleOrderOutput, err error) {
	// 获取客户信息
	getCustomerByIdOutput, err := service.Customer().GetCustomerById(ctx, pojo.GetCustomerByIdInput{CustomerId: in.CustomerId})
	if err != nil {
		return out, err
	}
	// 初始化订单信息
	currentTime := gtime.Now()
	out.SaleOrder = entity.SaleOrder{
		OrderNo:       s.generateOrderNo(consts.OrderTypeCGDD, currentTime),
		CustomerId:    getCustomerByIdOutput.CustomerId,
		CustomerName:  getCustomerByIdOutput.CustomerName,
		OrderQuantity: 0,
		OrderAmount:   0,
		CreateTime:    currentTime,
		CreateUser:    redis.Ctx(ctx).CheckLogin(),
		OrderStatus:   consts.OrderStatusInit,
	}
	// 插入并获取订单ID
	out.OrderId, err = dao.SaleOrder.Ctx(ctx).InsertAndGetId(out)
	return
}
