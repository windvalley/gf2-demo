// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package sdk

import (
	"context"

	"github.com/gogf/gf/contrib/sdk/httpclient/v2"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/windvalley/gf2-demo/api/demo"
	"github.com/windvalley/gf2-demo/api/demo/v1"
)

type implementerDemoV1 struct {
	*httpclient.Client
}

func (i *implementer) DemoV1() demo.IDemoV1 {
	var (
		client = httpclient.New(i.config)
		prefix = gstr.TrimRight(i.config.URL, "/") + ""
	)
	client.Client = client.Prefix(prefix)
	return &implementerDemoV1{client}
}

func (i *implementerDemoV1) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerDemoV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerDemoV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerDemoV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerDemoV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

