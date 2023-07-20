package demo

import (
	"context"

	v1 "gf2-demo/api/demo/v1"
	"gf2-demo/internal/model"
	"gf2-demo/internal/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	data := model.DemoCreateInput{
		Fielda: req.Fielda,
		Fieldb: req.Fieldb,
	}

	result, err := service.Demo().Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return &v1.CreateRes{ID: result.ID}, err
}
