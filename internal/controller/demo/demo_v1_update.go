package demo

import (
	"context"

	v1 "github.com/windvalley/gf2-demo/api/demo/v1"
	"github.com/windvalley/gf2-demo/internal/model"
	"github.com/windvalley/gf2-demo/internal/service"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	data := model.DemoUpdateInput{
		ID:     req.ID,
		Fielda: req.Fielda,
		Fieldb: req.Fieldb,
	}

	err = service.Demo().Update(ctx, data)

	return nil, err
}
