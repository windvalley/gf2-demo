package user

import (
	"context"

	"gf2-demo/internal/dao"
	"gf2-demo/internal/model"
	"gf2-demo/internal/model/do"
	"gf2-demo/internal/model/entity"
	"gf2-demo/internal/service"
)

type sDemo struct{}

func init() {
	service.RegisterDemo(New())
}

func New() *sDemo {
	return &sDemo{}
}

func (s *sDemo) Create(ctx context.Context, in model.DemoCreateInput) (*model.DemoCreateOutput, error) {
	id, err := dao.Demo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	return &model.DemoCreateOutput{
		ID: uint(id),
	}, nil
}

func (s *sDemo) GetInfo(ctx context.Context, fileda string) (*entity.Demo, error) {
	demoInfo, err := dao.Demo.Ctx(ctx).Where(do.Demo{
		Fielda: fileda,
	}).One()
	if err != nil {
		return nil, err
	}

	var demo *entity.Demo

	if err := demoInfo.Struct(&demo); err != nil {
		return nil, err
	}

	return demo, err
}
