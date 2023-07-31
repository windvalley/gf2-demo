package demo

import (
	"context"

	"github.com/windvalley/gf2-demo/api"
	v1 "github.com/windvalley/gf2-demo/api/demo/v1"
	"github.com/windvalley/gf2-demo/internal/model"
	"github.com/windvalley/gf2-demo/internal/service"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	result, err := service.Demo().List(ctx, model.DemoListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	list := result.List
	if len(list) == 0 {
		// 避免返回的空数组为null
		list = []model.DemoListOutputItem{}
	}

	return &v1.GetListRes{
		CommonPaginationRes: api.CommonPaginationRes{
			Total:    result.Total,
			PageNum:  result.PageNum,
			PageSize: result.PageSize,
		},
		List: list,
	}, nil
}
