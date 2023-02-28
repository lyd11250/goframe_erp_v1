package order

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"goframe-erp-v1/internal/consts"
	"goframe-erp-v1/internal/dao"
	"goframe-erp-v1/internal/model"
	"goframe-erp-v1/internal/model/entity"
	"goframe-erp-v1/internal/service"
)

type sOrder struct {
}

func New() *sOrder {
	return &sOrder{}
}

func init() {
	service.RegisterOrder(New())
}

func (s *sOrder) CreateOrder(ctx context.Context, in model.CreateOrderInput) (out model.CreateOrderOutput, err error) {
	// 初始化订单
	var order entity.Order
	// 初始化订单类型
	order.OrderType = in.OrderType
	// 初始化订单状态
	order.OrderStatus = consts.OrderStatusNew
	// 初始化订单时间
	order.OrderTime = gtime.Now()
	// 初始化订单号
	order.OrderNum = fmt.Sprintf("%v%v", consts.OrderPrefixMap[order.OrderType], order.OrderTime.Format("YmdHisu"))

	// 根据订单类型设置供应商或客户名称
	if in.OrderType == consts.OrderTypePurchase {
		supplier, err := service.Supplier().GetSupplierById(ctx, model.GetSupplierByIdInput{SupplierId: in.OrderSupplier})
		if err != nil {
			return model.CreateOrderOutput{}, err
		}
		if supplier.SupplierStatus != consts.StatusEnabled {
			return model.CreateOrderOutput{}, gerror.NewCode(gcode.CodeInvalidParameter, "供应商已禁用")
		}
		order.OrderSupplier = in.OrderSupplier
		order.OrderSupplierName = supplier.SupplierName
		order.OrderCustomer = -1
		order.OrderCustomerName = "SYSTEM"
	} else if in.OrderType == consts.OrderTypeSale {
		customer, err := service.Customer().GetCustomerById(ctx, model.GetCustomerByIdInput{CustomerId: in.OrderCustomer})
		if err != nil {
			return model.CreateOrderOutput{}, err
		}
		if customer.CustomerStatus != consts.StatusEnabled {
			return model.CreateOrderOutput{}, gerror.NewCode(gcode.CodeInvalidParameter, "客户已禁用")
		}
		order.OrderCustomer = in.OrderCustomer
		order.OrderCustomerName = customer.CustomerName
		order.OrderSupplier = -1
		order.OrderSupplierName = "SYSTEM"
	}

	// 储存并获取ID
	orderId, err := dao.Order.Ctx(ctx).InsertAndGetId(order)
	if err != nil {
		return model.CreateOrderOutput{}, err
	}

	out.OrderId = orderId
	out.OrderNum = order.OrderNum

	return
}

func (s *sOrder) DeleteOrder(ctx context.Context, in model.DeleteOrderInput) (err error) {
	_, err = dao.Order.Ctx(ctx).WherePri(in.OrderId).Delete()
	if err != nil {
		return err
	}
	_, err = dao.OrderItem.Ctx(ctx).Where(dao.OrderItem.Columns().OrderId, in.OrderId).Delete()
	return
}

func (s *sOrder) GetOrderList(ctx context.Context, in model.GetOrderListInput) (out model.GetOrderListOutput, err error) {
	// 初始化分页条件
	var limit int
	var offset int
	// 初始化分页条件
	if in.PageSize > 0 {
		limit = in.PageSize
		offset = (in.Page - 1) * in.PageSize
	}

	// 初始化查询条件
	where := gmap.NewStrAnyMap()
	where.Set(dao.Order.Columns().OrderType, in.OrderType)
	if in.OrderNum != "" {
		where.Set(dao.Order.Columns().OrderNum, in.OrderNum)
	}

	// 查询列表
	err = dao.Order.Ctx(ctx).
		Where(where.Map()).
		OrderDesc(dao.Order.Columns().OrderNum).
		Limit(limit).
		Offset(offset).Scan(&out.List)
	if err != nil {
		return model.GetOrderListOutput{}, err
	}

	// 查询总数
	count, err := dao.Order.Ctx(ctx).Count(where.Map())
	if err != nil {
		return model.GetOrderListOutput{}, err
	}
	out.Total = count
	out.Pages = count / limit
	if count%limit > 0 {
		out.Pages++
	}

	return
}

func (s *sOrder) GetOrderById(ctx context.Context, in model.GetOrderByIdInput) (out model.GetOrderByIdOutput, err error) {
	err = dao.Order.Ctx(ctx).WherePri(in.OrderId).Scan(&out.Order)
	if err != nil {
		return model.GetOrderByIdOutput{}, err
	}
	err = dao.OrderItem.Ctx(ctx).Where(dao.OrderItem.Columns().OrderId, in.OrderId).Scan(&out.Items)
	return
}

func (s *sOrder) SetOrderItem(ctx context.Context, in model.SetOrderItemInput) (err error) {
	var order entity.Order
	err = dao.Order.Ctx(ctx).WherePri(in.OrderId).Scan(&order)
	if err != nil {
		return err
	}
	// 检查每个订单项
	for i := range in.Items {
		// 检查订单ID
		in.Items[i].OrderId = in.OrderId

		// 检查商品
		goodsOutput, err := service.Goods().GetGoodsById(ctx, model.GetGoodsByIdInput{GoodsId: in.Items[i].GoodsId})
		if err != nil {
			return err
		}
		if goodsOutput.GoodsStatus != consts.StatusEnabled {
			return gerror.NewCodef(gcode.CodeInvalidParameter, "商品(%v)已禁用", goodsOutput.GoodsName)
		}
		in.Items[i].GoodsName = goodsOutput.GoodsName

		// 检查价格
		if in.Items[i].GoodsPrice != goodsOutput.GoodsPrice {
			return gerror.NewCodef(gcode.CodeInvalidParameter, "商品(%v)价格被非法修改", goodsOutput.GoodsName)
		}
		in.Items[i].OrderItemAmount = in.Items[i].GoodsPrice * float64(in.Items[i].Quantity)

		// 加入订单总额
		order.OrderAmount += in.Items[i].OrderItemAmount

		// 针对销售单检查库存
		if order.OrderType == consts.OrderTypeSale {

			checkInventoryOutput, err := service.Inventory().
				CheckInventory(ctx,
					model.CheckInventoryInput{
						GoodsId:  in.Items[i].GoodsId,
						Quantity: int(in.Items[i].Quantity),
					})
			if err != nil {
				return err
			}
			if !checkInventoryOutput.Enough {
				return gerror.NewCodef(gcode.CodeInvalidParameter, "商品(%v)库存不足", goodsOutput.GoodsName)
			}
		}
	}
	// 更新订单状态
	order.OrderStatus = consts.OrderStatusProcessing
	// 插入数据
	_, err = dao.OrderItem.Ctx(ctx).Insert(in.Items)
	if err != nil {
		return err
	}
	_, err = dao.Order.Ctx(ctx).WherePri(in.OrderId).Update(order)
	return
}
