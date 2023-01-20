package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "goframe-erp-v1/api/v1"
	"goframe-erp-v1/internal/model"
	"goframe-erp-v1/internal/service"
)

type cAccess struct {
}

var Access cAccess

func (c *cAccess) AddAccess(ctx context.Context, req *v1.AddAccessReq) (res *v1.AddAccessRes, err error) {
	output, err := service.Access().AddAccess(ctx, model.AddAccessInput{
		AccessTitle: req.AccessTitle,
		AccessUri:   req.AccessUri,
	})
	if err != nil {
		return nil, err
	}
	res = &v1.AddAccessRes{AccessId: output.AccessId}
	return
}

func (c *cAccess) UpdateAccess(ctx context.Context, req *v1.UpdateAccessReq) (res *v1.UpdateAccessRes, err error) {
	input := model.UpdateAccessInput{}
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	err = service.Access().UpdateAccess(ctx, input)
	return
}

func (c *cAccess) DeleteAccess(ctx context.Context, req *v1.DeleteAccessReq) (res *v1.DeleteAccessRes, err error) {
	input := model.DeleteAccessInput{}
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	err = service.Access().DeleteAccess(ctx, input)
	return
}
