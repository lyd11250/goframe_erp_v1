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

func (s *sOrder) GetPurchaseOrder(ctx context.Context, in order.GetPurchaseOrderInput) (out order.GetPurchaseOrderOutput, err error) {
	orderResult, err := dao.PurchaseOrder.Ctx(ctx).OmitNil().One(in)

	if err != nil {
		return order.GetPurchaseOrderOutput{}, err
	}
	if orderResult.IsEmpty() {
		return order.GetPurchaseOrderOutput{}, gerror.NewCode(gcode.CodeNotFound, "订单不存在")
	}
	err = orderResult.Struct(&out.Order)
	if err != nil {
		return order.GetPurchaseOrderOutput{}, err
	}
	err = dao.OrderItem.Ctx(ctx).Where(dao.OrderItem.Columns().OrderId, out.Order.OrderId).Scan(&out.Items)
	return
}

func (s *sOrder) GetPurchaseOrderList(ctx context.Context, in order.GetPurchaseOrderListInput) (out order.GetPurchaseOrderListOutput, err error) {
	err = dao.PurchaseOrder.Ctx(ctx).Page(in.Page, in.PageSize).Scan(&out.List)
	if err != nil {
		return order.GetPurchaseOrderListOutput{}, err
	}
	out.Total, err = dao.PurchaseOrder.Ctx(ctx).Count()
	if err != nil {
		return order.GetPurchaseOrderListOutput{}, err
	}
	out.Pages = out.Total / in.PageSize
	if out.Total%in.PageSize != 0 {
		out.Pages++
	}
	return
}

func (s *sOrder) CreatePurchaseOrder(ctx context.Context, in order.CreatePurchaseOrderInput) (out order.CreatePurchaseOrderOutput, err error) {
	// 获取供应商信息
	getSupplierByIdOutput, err := service.Supplier().GetSupplierById(ctx, pojo.GetSupplierByIdInput{SupplierId: in.SupplierId})
	if err != nil {
		return out, err
	}
	// 初始化订单信息
	currentTime := gtime.Now()
	out.PurchaseOrder = entity.PurchaseOrder{
		OrderNo:       s.generateOrderNo(consts.OrderTypeCGDD, currentTime),
		SupplierId:    getSupplierByIdOutput.SupplierId,
		SupplierName:  getSupplierByIdOutput.SupplierName,
		OrderQuantity: 0,
		OrderAmount:   0,
		CreateTime:    currentTime,
		CreateUser:    redis.Ctx(ctx).CheckLogin(),
		OrderStatus:   consts.OrderStatusInit,
	}
	// 插入并获取订单ID
	out.OrderId, err = dao.PurchaseOrder.Ctx(ctx).InsertAndGetId(out)
	return
}
