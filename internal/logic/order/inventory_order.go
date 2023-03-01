package order

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-erp-v1/internal/consts"
	"goframe-erp-v1/internal/dao"
	"goframe-erp-v1/internal/model/entity"
	"goframe-erp-v1/internal/model/pojo/order"
	"goframe-erp-v1/utility/redis"
)

func (s *sOrder) GetInventoryOrder(ctx context.Context, in order.GetInventoryOrderInput) (out order.GetInventoryOrderOutput, err error) {
	// 获取订单信息
	orderResult, err := dao.InventoryOrder.Ctx(ctx).OmitNil().One(in)
	if err != nil {
		return out, err
	}
	if orderResult.IsEmpty() {
		return out, gerror.NewCode(gcode.CodeNotFound, "订单不存在")
	}
	err = orderResult.Struct(&out.Order)
	if err != nil {
		return out, err
	}
	// 获取订单项信息
	err = dao.OrderItem.Ctx(ctx).
		OmitNil().
		Where(dao.OrderItem.Columns().OrderId, out.Order.OrderId).
		Scan(&out.Items)
	return
}

func (s *sOrder) GetInventoryOrderList(ctx context.Context, in order.GetInventoryOrderListInput) (out order.GetInventoryOrderListOutput, err error) {
	err = dao.InventoryOrder.Ctx(ctx).
		Page(in.Page, in.PageSize).
		Where(dao.InventoryOrder.Columns().OrderType, in.OrderType).
		Scan(&out.List)
	if err != nil {
		return order.GetInventoryOrderListOutput{}, err
	}
	out.Total, err = dao.InventoryOrder.Ctx(ctx).Where(dao.InventoryOrder.Columns().OrderType, in.OrderType).Count()
	if err != nil {
		return order.GetInventoryOrderListOutput{}, err
	}
	out.Pages = out.Total / in.PageSize
	if out.Total%in.PageSize != 0 {
		out.Pages++
	}
	return
}

func (s *sOrder) CreateInventoryOrder(ctx context.Context, in order.CreateInventoryOrderInput) (out order.CreateInventoryOrderOutput, err error) {
	// 获取源订单信息
	if in.OrderType == consts.OrderTypeXSCK {
		output, err := s.GetSaleOrder(ctx, order.GetSaleOrderInput{OrderId: &in.POrderId})
		if err != nil {
			return out, err
		}
		out.POrder = gconv.MapDeep(output)
	} else if in.OrderType == consts.OrderTypeCGRK {
		// 获取采购订单信息
		output, err := s.GetPurchaseOrder(ctx, order.GetPurchaseOrderInput{OrderId: &in.POrderId})
		if err != nil {
			return out, err
		}
		out.POrder = gconv.MapDeep(output)
	}
	// 初始化订单信息
	var currentTime = gtime.Now()
	out.Order = entity.InventoryOrder{
		OrderNo:     s.generateOrderNo(in.OrderType, currentTime),
		OrderType:   in.OrderType,
		POrderId:    in.POrderId,
		CreateTime:  currentTime,
		CreateUser:  redis.Ctx(ctx).CheckLogin(),
		OrderStatus: consts.OrderStatusInit,
	}
	// 插入并获取订单ID
	out.Order.OrderId, err = dao.InventoryOrder.Ctx(ctx).InsertAndGetId(out.Order)

	return
}

func (s *sOrder) CancelCreateInventoryOrder(ctx context.Context, in order.CancelCreateInventoryOrderInput) (err error) {
	// 获取订单信息
	orderInfo, err := s.GetInventoryOrder(ctx, order.GetInventoryOrderInput{OrderId: in.OrderId, OrderNo: in.OrderNo})
	if err != nil {
		return err
	}
	// 判断订单状态
	if orderInfo.Order.OrderStatus != consts.OrderStatusInit {
		return gerror.NewCode(gcode.CodeInvalidParameter, "订单状态错误")
	}
	// 删除订单
	_, err = dao.InventoryOrder.Ctx(ctx).OmitEmpty().Delete(in)
	if err != nil {
		return err
	}
	// 删除订单项
	_, err = dao.OrderItem.Ctx(ctx).
		OmitEmpty().
		Where(dao.OrderItem.Columns().OrderId, orderInfo.Order.OrderId).
		Delete()
	return
}
