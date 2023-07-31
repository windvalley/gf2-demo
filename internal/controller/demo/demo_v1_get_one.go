package demo

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/windvalley/gf2-demo/api/demo/v1"
	"github.com/windvalley/gf2-demo/internal/codes"
	"github.com/windvalley/gf2-demo/internal/service"
)

func (c *ControllerV1) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	demoInfo, err := service.Demo().Get(ctx, req.Fielda)
	if err != nil {
		return nil, err
	}

	if demoInfo == nil {
		return nil, gerror.WrapCode(codes.CodeNotFound, fmt.Errorf("fielda '%s'", req.Fielda))
	}

	return &v1.GetOneRes{
		ID:        demoInfo.Id,
		Fielda:    demoInfo.Fielda,
		CreatedAt: demoInfo.CreatedAt,
		UpdatedAt: demoInfo.UpdatedAt,
	}, nil
}
