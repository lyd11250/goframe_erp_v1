package inventory

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-erp-v1/internal/dao"
	"goframe-erp-v1/internal/model"
	"goframe-erp-v1/internal/model/entity"
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

func (s *sInventory) CheckInventory(ctx context.Context, in model.CheckInventoryInput) (out model.CheckInventoryOutput, err error) {
	quantity, err := dao.Inventory.Ctx(ctx).
		Where(dao.Inventory.Columns().GoodsId, in.GoodsId).
		Sum(dao.Inventory.Columns().Quantity)
	if err != nil {
		return model.CheckInventoryOutput{}, err
	}
	out.Enough = quantity >= float64(in.Quantity)
	return
}

func (s *sInventory) DeleteInventory(ctx context.Context, in model.DeleteInventoryInput) (err error) {
	_, err = dao.Inventory.Ctx(ctx).Where(in).Delete()
	return
}

func (s *sInventory) AddInventory(ctx context.Context, in model.AddInventoryInput) (err error) {
	// 检查商品是否停用
	goodsEnabledOutput, err := service.Goods().CheckGoodsEnabled(ctx, model.CheckGoodsEnabledInput{GoodsId: in.GoodsId})
	if err != nil {
		return err
	}
	if !goodsEnabledOutput.Enabled {
		return gerror.NewCode(gcode.CodeInvalidParameter, "商品已停用")
	}

	// 检查库存是否存在
	currentInventoryOutput, err := dao.Inventory.Ctx(ctx).
		Where(g.Map{
			dao.Inventory.Columns().GoodsId:   in.GoodsId,
			dao.Inventory.Columns().GoodsCost: in.GoodsCost,
		}).One()
	if err != nil {
		return err
	}

	// 不存在则新增
	if currentInventoryOutput.IsEmpty() {
		_, err = dao.Inventory.Ctx(ctx).Insert(in)
		return
	}

	// 存在则更新
	var currentInventory entity.Inventory
	err = currentInventoryOutput.Struct(&currentInventory)
	if err != nil {
		return err
	}
	currentInventory.Quantity += in.Quantity
	err = s.UpdateInventory(ctx, model.UpdateInventoryInput{Inventory: currentInventory})
	return
}

func (s *sInventory) UpdateInventory(ctx context.Context, in model.UpdateInventoryInput) (err error) {
	// 检查库存数量是否小于0
	if in.Quantity < 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "库存数量不能小于0")
	}
	// 检查库存数量是否等于0，等于0则删除
	if in.Quantity == 0 {
		err = s.DeleteInventory(ctx, model.DeleteInventoryInput{
			GoodsId:   in.GoodsId,
			GoodsCost: in.GoodsCost,
		})
		return
	}
	_, err = dao.Inventory.Ctx(ctx).
		Data(dao.Inventory.Columns().Quantity, in.Quantity).
		Where(g.Map{
			dao.Inventory.Columns().GoodsId:   in.GoodsId,
			dao.Inventory.Columns().GoodsCost: in.GoodsCost,
		}).
		Update()
	return
}

func (s *sInventory) ReduceInventory(ctx context.Context, in model.ReduceInventoryInput) (err error) {
	// 检查商品是否停用
	goodsEnabledOutput, err := service.Goods().CheckGoodsEnabled(ctx, model.CheckGoodsEnabledInput{GoodsId: in.GoodsId})
	if err != nil {
		return err
	}
	if !goodsEnabledOutput.Enabled {
		return gerror.NewCode(gcode.CodeInvalidParameter, "商品已停用")
	}

	// 检查库存是否存在
	currentInventoriesOutput, err := dao.Inventory.Ctx(ctx).
		Where(dao.Inventory.Columns().GoodsId, in.GoodsId).
		OrderAsc(dao.Inventory.Columns().GoodsCost).
		All()
	if err != nil {
		return err
	}
	if currentInventoriesOutput.IsEmpty() {
		return gerror.NewCode(gcode.CodeInvalidParameter, "库存不足")
	}

	// 检查库存是否足够
	checkInventoryOutput, err := s.CheckInventory(ctx, model.CheckInventoryInput{
		GoodsId:  in.GoodsId,
		Quantity: in.Quantity,
	})
	if err != nil {
		return err
	}
	if !checkInventoryOutput.Enough {
		return gerror.NewCode(gcode.CodeInvalidParameter, "库存不足")
	}

	// 减库存
	var currentInventories []entity.Inventory
	err = currentInventoriesOutput.Structs(&currentInventories)
	if err != nil {
		return err
	}
	quantity := in.Quantity
	for i := range currentInventories {
		if quantity <= 0 {
			break
		}
		if currentInventories[i].Quantity >= quantity {
			currentInventories[i].Quantity -= quantity
			quantity = 0
		} else {
			quantity -= currentInventories[i].Quantity
			currentInventories[i].Quantity = 0
		}
	}

	// 更新库存事务
	err = dao.Inventory.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (e error) {
		for i := range currentInventories {
			_, e = dao.Inventory.Ctx(ctx).TX(tx).
				Data(dao.Inventory.Columns().Quantity, currentInventories[i].Quantity).
				Where(g.Map{
					dao.Inventory.Columns().GoodsId:   currentInventories[i].GoodsId,
					dao.Inventory.Columns().GoodsCost: currentInventories[i].GoodsCost,
				}).
				Update()
			if e != nil {
				return
			}
		}
		return
	})
	return
}

func (s *sInventory) GetGoodsInventory(ctx context.Context, in model.GetGoodsInventoryInput) (out model.GetGoodsInventoryOutput, err error) {
	err = dao.Inventory.Ctx(ctx).
		Where(dao.Inventory.Columns().GoodsId, in.GoodsId).
		OrderAsc(dao.Inventory.Columns().GoodsCost).
		Scan(&out.List)
	if err != nil {
		return
	}
	for _, inventory := range out.List {
		out.Sum += inventory.Quantity
		out.Amount += inventory.GoodsCost * float64(inventory.Quantity)
	}

	return
}
