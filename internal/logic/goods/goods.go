package goods

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-erp-v1/internal/consts"
	"goframe-erp-v1/internal/dao"
	"goframe-erp-v1/internal/model"
	"goframe-erp-v1/internal/service"
)

type sGoods struct {
}

func New() *sGoods {
	return &sGoods{}
}

func init() {
	service.RegisterGoods(New())
}

func (s *sGoods) GetGoodsById(ctx context.Context, in model.GetGoodsByIdInput) (out model.GetGoodsByIdOutput, err error) {
	result, err := dao.Goods.Ctx(ctx).WherePri(in.GoodsId).One()
	if err != nil {
		return model.GetGoodsByIdOutput{}, err
	}
	if result.IsEmpty() {
		return out, gerror.NewCode(gcode.CodeNotFound, "商品不存在")
	}
	err = result.Struct(&out)
	return
}

func (s *sGoods) GetGoodsByName(ctx context.Context, in model.GetGoodsByNameInput) (out model.GetGoodsByNameOutput, err error) {
	result, err := dao.Goods.Ctx(ctx).
		WhereLike(dao.Goods.Columns().GoodsName, "%"+*in.GoodsName+"%").
		OrderDesc(dao.Goods.Columns().GoodsStatus).
		All()
	if err != nil {
		return out, err
	}
	if result.IsEmpty() {
		return out, gerror.NewCode(gcode.CodeNotFound, "商品不存在")
	}
	err = result.Structs(&out.List)
	return
}

func (s *sGoods) AddGoods(ctx context.Context, in model.AddGoodsInput) (out model.AddGoodsOutput, err error) {
	id, err := dao.Goods.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return model.AddGoodsOutput{}, err
	}
	out.GoodsId = id
	return
}

func (s *sGoods) UpdateGoods(ctx context.Context, in model.UpdateGoodsInput) (err error) {
	_, err = dao.Goods.Ctx(ctx).OmitNil().Data(in).WherePri(in.GoodsId).Update()
	return
}

func (s *sGoods) GetGoodsUnits(ctx context.Context) (out model.GetGoodsUnitsOutput, err error) {
	column := dao.Goods.Columns().GoodsUnit
	result, err := dao.Goods.Ctx(ctx).Fields(column).Group(column).All()
	if err != nil {
		return model.GetGoodsUnitsOutput{}, err
	}
	for _, v := range result.Array() {
		out.List = append(out.List, v.String())
	}
	return
}

func (s *sGoods) GetGoodsSuppliers(ctx context.Context, in model.GetGoodsSuppliersInput) (out model.GetGoodsSuppliersOutput, err error) {
	err = dao.GoodsSupplierRel.Ctx(ctx).
		Fields(dao.GoodsSupplierRel.Columns().SupplierId, dao.GoodsSupplierRel.Columns().SupplyPrice).
		OrderAsc(dao.GoodsSupplierRel.Columns().SupplyPrice).
		Scan(&out.List, dao.GoodsSupplierRel.Columns().GoodsId, in.GoodsId)
	return
}

func (s *sGoods) AddGoodsSupplier(ctx context.Context, in model.AddGoodsSupplierInput) (err error) {
	// 判断供应商是否停用
	supplier, err := service.Supplier().GetSupplierById(ctx, model.GetSupplierByIdInput{SupplierId: in.SupplierId})
	if err != nil {
		return err
	}
	if supplier.SupplierStatus == consts.StatusDisabled {
		return gerror.NewCode(gcode.CodeInvalidParameter, "该供应商已停用")
	}

	// 判断供应商是否已存在
	count, err := dao.GoodsSupplierRel.Ctx(ctx).Count(g.Map{
		dao.GoodsSupplierRel.Columns().GoodsId:    in.GoodsId,
		dao.GoodsSupplierRel.Columns().SupplierId: in.SupplierId,
	})
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "该供应商已存在")
	}

	_, err = dao.GoodsSupplierRel.Ctx(ctx).Insert(in)
	return
}

func (s *sGoods) UpdateGoodsSupplier(ctx context.Context, in model.UpdateGoodsSupplierInput) (err error) {
	_, err = dao.GoodsSupplierRel.Ctx(ctx).
		Data(dao.GoodsSupplierRel.Columns().SupplyPrice, in.SupplyPrice).
		Where(g.Map{
			dao.GoodsSupplierRel.Columns().GoodsId:    in.GoodsId,
			dao.GoodsSupplierRel.Columns().SupplierId: in.SupplierId,
		}).
		Update()
	return
}

func (s *sGoods) DeleteGoodsSupplier(ctx context.Context, in model.DeleteGoodsSupplierInput) (err error) {
	_, err = dao.GoodsSupplierRel.Ctx(ctx).Delete(in)
	return
}

func (s *sGoods) CheckGoodsEnabled(ctx context.Context, in model.CheckGoodsEnabledInput) (out model.CheckGoodsEnabledOutput, err error) {
	result, err := dao.Goods.Ctx(ctx).Fields(dao.Goods.Columns().GoodsStatus).WherePri(in.GoodsId).One()
	if err != nil {
		return model.CheckGoodsEnabledOutput{}, err
	}
	if result.IsEmpty() {
		return out, gerror.NewCode(gcode.CodeNotFound, "商品不存在")
	}
	out.Enabled = result.Map()["goods_status"] == consts.StatusEnabled
	return
}
