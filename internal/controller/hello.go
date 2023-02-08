package controller

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	v1 "gf2-demo/api/v1"
	"gf2-demo/internal/codes"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	// res = &v1.HelloRes{Name: "World"}

	err = errors.New("some error")

	err = gerror.WrapCode(codes.CodeNotAuthorized, err)

	// panic("failed")

	g.Log().Error(ctx, "failed")

	return
}
