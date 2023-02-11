package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 接口文档: https://goframe.org/pages/viewpage.action?pageId=47703679
// 请求输入: https://goframe.org/pages/viewpage.action?pageId=1114483
type DemoInfoReq struct {
	g.Meta `path:"/demo/info" method:"get" tags:"DemoService" summary:"Get demo info by field1"`
	Fielda string `p:"filed1"  v:"required|length:4,30"`
}

type DemoInfoRes struct {
	ID        uint        `json:"id"`
	Fielda    string      `json:"fielda"`
	CreatedAt *gtime.Time `json:"created_at"`
	UpdatedAt *gtime.Time `json:"updated_at"`
}

type DemoCreateReq struct {
	g.Meta `path:"/demo" method:"post" tags:"DemoService" summary:"Create a demo record"`
	Fielda string `p:"fileda"  v:"required|length:4,30"`
	Fieldb string `p:"filedb"  v:"required|length:10,30"`
}

type DemoCreateRes struct {
	ID uint `json:"id"`
}
