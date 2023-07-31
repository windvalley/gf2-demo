package demo

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/windvalley/gf2-demo/internal/codes"
	"github.com/windvalley/gf2-demo/internal/dao"
	"github.com/windvalley/gf2-demo/internal/model"
	"github.com/windvalley/gf2-demo/internal/model/do"
	"github.com/windvalley/gf2-demo/internal/model/entity"
	"github.com/windvalley/gf2-demo/internal/service"
)

type sDemo struct{}

func init() {
	service.RegisterDemo(New())
}

func New() *sDemo {
	return &sDemo{}
}

func (s *sDemo) Create(ctx context.Context, in model.DemoCreateInput) (*model.DemoCreateOutput, error) {
	notFound, err := s.FieldaNotFound(ctx, in.Fielda)
	if err != nil {
		return nil, err
	}
	if !notFound {
		err1 := gerror.WrapCode(codes.CodeNotAvailable, fmt.Errorf("fielda '%s' already exists", in.Fielda))
		return nil, err1
	}

	id, err := dao.Demo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	return &model.DemoCreateOutput{
		ID: uint(id),
	}, nil
}

func (s *sDemo) Update(ctx context.Context, in model.DemoUpdateInput) error {
	idNotFound, err := s.IDNotFound(ctx, in.ID)
	if err != nil {
		return err
	}
	if idNotFound {
		return gerror.WrapCode(codes.CodeNotFound, fmt.Errorf("id '%d' not found", in.ID))
	}

	count, err := dao.Demo.Ctx(ctx).Where("id!=?", in.ID).Where("fielda=?", in.Fielda).Count()
	if err != nil {
		return err
	}
	if count != 0 {
		return gerror.WrapCode(codes.CodeNotAvailable, fmt.Errorf("fielda '%s' already exists", in.Fielda))
	}

	_, err1 := dao.Demo.Ctx(ctx).OmitEmpty().Data(in).Where(do.Demo{
		Id: in.ID,
	}).Update()

	return err1
}

func (s *sDemo) Get(ctx context.Context, fileda string) (*entity.Demo, error) {
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

func (s *sDemo) Delete(ctx context.Context, id uint) error {
	res, err := dao.Demo.Ctx(ctx).Where(do.Demo{
		Id: id,
	}).Delete()
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n == 0 {
		return gerror.NewCode(codes.CodeNotFound)
	}

	return nil
}

func (s *sDemo) List(ctx context.Context, in model.DemoListInput) (*model.DemoListOutput, error) {
	m := dao.Demo.Ctx(ctx)

	out := &model.DemoListOutput{
		PageNum:  in.PageNum,
		PageSize: in.PageSize,
	}

	listModel := m.Page(in.PageNum, in.PageSize)

	// 按照更新时间排序
	listModel = listModel.OrderDesc(dao.Demo.Columns().UpdatedAt)

	if err := listModel.Scan(&out.List); err != nil {
		return nil, err
	}
	if len(out.List) == 0 {
		return &model.DemoListOutput{}, nil
	}

	count, err := m.Count()
	if err != nil {
		return nil, err
	}
	out.Total = count

	return out, nil
}

func (s *sDemo) IDNotFound(ctx context.Context, id uint) (bool, error) {
	count, err := dao.Demo.Ctx(ctx).Where(do.Demo{
		Id: id,
	}).Count()
	if err != nil {
		return false, err
	}

	return count == 0, nil
}

func (s *sDemo) FieldaNotFound(ctx context.Context, fielda string) (bool, error) {
	count, err := dao.Demo.Ctx(ctx).Where(do.Demo{
		Fielda: fielda,
	}).Count()
	if err != nil {
		return false, err
	}

	return count == 0, nil
}
