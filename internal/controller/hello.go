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
	// 正常响应测试
	// res = &v1.HelloRes{Name: "World"}

	// 调用其他HTTP服务测试
	// foo := g.Client().GetContent(ctx, "http://localhost:8000/v1/hello")

	// 通用日志打印测试
	g.Log().Info(ctx, "hello world")
	g.Log().Errorf(ctx, "xxx failed")

	// panic错误测试
	// panic("panic test")

	// 自定义错误测试
	err = errors.New("some error")
	err = gerror.WrapCode(codes.CodeNotAuthorized, err)

	return
}
