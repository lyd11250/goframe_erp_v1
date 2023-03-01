package inventory

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-erp-v1/internal/dao"
	"goframe-erp-v1/internal/model/entity"
	"goframe-erp-v1/internal/model/pojo"
	"goframe-erp-v1/internal/service"
)

type sInventory struct {
}

func New() *sInventory {
	return &sInventory{}
}

func init() {
	service.RegisterInventory(New())
}

func (s *sInventory) GetGoodsInventory(ctx context.Context, in pojo.GetGoodsInventoryInput) (out pojo.GetGoodsInventoryOutput, err error) {
	_, err = service.Goods().GetGoodsById(ctx, pojo.GetGoodsByIdInput{GoodsId: in.GoodsId})
	if err != nil {
		return out, err
	}
	result, err := dao.Inventory.Ctx(ctx).WherePri(in.GoodsId).One()
	if err != nil {
		return out, err
	}
	if result.IsEmpty() {
		out = pojo.GetGoodsInventoryOutput{
			Inventory: entity.Inventory{
				GoodsId:  in.GoodsId,
				Quantity: 0,
				Price:    0,
				Amount:   0,
			},
		}
		return
	}
	err = gconv.Struct(result, &out)
	return
}

func (s *sInventory) GetInventoryList(ctx context.Context, in pojo.GetInventoryListInput) (out pojo.GetInventoryListOutput, err error) {
	err = dao.Inventory.Ctx(ctx).Page(in.Page, in.PageSize).Scan(&out.List)
	if err != nil {
		return
	}
	out.Total, err = dao.Inventory.Ctx(ctx).Count()
	if err != nil {
		return
	}
	out.Pages = out.Total / in.PageSize
	if out.Total%in.PageSize > 0 {
		out.Pages++
	}
	return
}

func (s *sInventory) AddInventory(ctx context.Context, in pojo.AddInventoryInput) (out pojo.AddInventoryOutput, err error) {
	// 检查输入
	if in.Quantity <= 0 || in.Price < 0 {
		return out, gerror.NewCode(gcode.CodeInvalidParameter, "数量或单价不能小于0")
	}
	if float64(in.Quantity)*in.Price != in.Amount {
		return out, gerror.NewCode(gcode.CodeInvalidParameter, "金额不正确")
	}
	// 加权平均法
	// 1. 获取当前库存
	currentInventory, err := s.GetGoodsInventory(ctx, pojo.GetGoodsInventoryInput{GoodsId: in.GoodsId})
	if err != nil {
		return
	}
	// 若当前库存为0，则直接更新库存
	if currentInventory.Quantity == 0 {
		_, err = dao.Inventory.Ctx(ctx).Insert(in)
		if err != nil {
			return out, err
		}
	}
	// 2. 若当前库存不为0，则计算新的库存
	out.Before = currentInventory.Inventory
	// 3. 计算新的库存
	currentInventory.Quantity += in.Quantity
	currentInventory.Amount += in.Amount
	currentInventory.Price = currentInventory.Amount / float64(currentInventory.Quantity)
	// 4. 更新库存
	_, err = dao.Inventory.Ctx(ctx).WherePri(in.GoodsId).Data(currentInventory).Update()
	if err != nil {
		return
	}
	out.After = currentInventory.Inventory
	return
}

func (s *sInventory) ReduceInventory(ctx context.Context, in pojo.ReduceInventoryInput) (out pojo.ReduceInventoryOutput, err error) {
	// 获取当前库存
	currentInventory, err := s.GetGoodsInventory(ctx, pojo.GetGoodsInventoryInput{GoodsId: in.GoodsId})
	if err != nil {
		return
	}
	// 检查库存是否充足
	if currentInventory.Quantity < in.Quantity {
		return out, gerror.NewCode(gcode.CodeInvalidParameter, "库存不足")
	}
	// 更新返回参数
	out.Before = currentInventory.Inventory
	// 更新库存
	currentInventory.Quantity -= in.Quantity
	currentInventory.Amount = float64(currentInventory.Quantity) * currentInventory.Price
	_, err = dao.Inventory.Ctx(ctx).WherePri(in.GoodsId).Data(currentInventory).Update()
	if err != nil {
		return
	}
	out.After = currentInventory.Inventory
	return
}
