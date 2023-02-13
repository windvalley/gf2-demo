package controller

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "gf2-demo/api/v1"
	"gf2-demo/internal/codes"
	"gf2-demo/internal/model"
	"gf2-demo/internal/service"
)

var (
	Demo = cDemo{}
)

type cDemo struct{}

func (c *cDemo) Info(ctx context.Context, req *v1.DemoInfoReq) (*v1.DemoInfoRes, error) {
	// 调用其他HTTP服务测试
	// foo := g.Client().GetContent(ctx, "http://localhost:8000/v1/hello")

	// 通用日志打印测试
	// g.Log().Info(ctx, "hello world")
	// g.Log().Errorf(ctx, "xxx failed")

	// panic错误测试
	// panic("panic test")

	// 自定义错误测试
	// err = errors.New("no permission")
	// err = gerror.WrapCode(codes.CodeNotAuthorized, err)

	demoInfo, err := service.Demo().GetInfo(ctx, req.Fielda)
	if err != nil {
		return nil, err
	}

	if demoInfo == nil {
		return nil, gerror.WrapCode(codes.CodeNotFound, fmt.Errorf("fielda '%s'", req.Fielda))
	}

	return &v1.DemoInfoRes{
		ID:        demoInfo.Id,
		Fielda:    demoInfo.Fielda,
		CreatedAt: demoInfo.CreatedAt,
		UpdatedAt: demoInfo.UpdatedAt,
	}, nil
}

func (c *cDemo) Create(ctx context.Context, req *v1.DemoCreateReq) (*v1.DemoCreateRes, error) {
	data := model.DemoCreateInput{
		Fielda: req.Fielda,
		Fieldb: req.Fieldb,
	}

	res, err := service.Demo().Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return &v1.DemoCreateRes{ID: res.ID}, err
}

func (c *cDemo) Update(ctx context.Context, req *v1.DemoUpdateReq) (*v1.DemoUpdateRes, error) {
	data := model.DemoUpdateInput{
		ID:     req.ID,
		Fielda: req.Fielda,
		Fieldb: req.Fieldb,
	}

	err := service.Demo().Update(ctx, data)

	return nil, err
}

func (c *cDemo) Delete(ctx context.Context, req *v1.DemoDeleteReq) (*v1.DemoDeleteRes, error) {
	err := service.Demo().Delete(ctx, req.ID)

	return nil, err
}
