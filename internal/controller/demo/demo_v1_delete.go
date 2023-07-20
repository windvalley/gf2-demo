package demo

import (
	"context"

	v1 "gf2-demo/api/demo/v1"
	"gf2-demo/internal/service"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.Demo().Delete(ctx, req.ID)

	return nil, err
}
