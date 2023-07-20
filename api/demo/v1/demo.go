package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gf2-demo/api"
	"gf2-demo/internal/model"
)

/*
API文档: https://goframe.org/pages/viewpage.action?pageId=47703679
路由规则: https://goframe.org/pages/viewpage.action?pageId=1114257
请求输入: https://goframe.org/pages/viewpage.action?pageId=1114483
字段校验: https://goframe.org/pages/viewpage.action?pageId=1114367
*/

// 查询单条记录示例.
type GetOneReq struct {
	g.Meta `path:"/demo/:fielda" method:"get" tags:"DemoService" summary:"Get demo info by fielda"`
	Fielda string `p:"fielda" in:"path" v:"required|length:4,30"`
}

type GetOneRes struct {
	ID        uint        `json:"id"`
	Fielda    string      `json:"fielda"`
	CreatedAt *gtime.Time `json:"created_at"`
	UpdatedAt *gtime.Time `json:"updated_at"`
}

// 查询记录列表示例.
type GetListReq struct {
	g.Meta `path:"/demo" method:"get" tags:"DemoService" summary:"Get demo records list"`
	api.CommonPaginationReq
}

type GetListRes struct {
	api.CommonPaginationRes

	List []model.DemoListOutputItem `json:"list"`
}

// 创建记录示例.
type CreateReq struct {
	g.Meta `path:"/demo" method:"post" tags:"DemoService" summary:"Create a demo record"`
	Fielda string `p:"fileda" v:"required|passport|length:4,30"`
	Fieldb string `p:"filedb" v:"required|length:10,30"`
}

type CreateRes struct {
	ID uint `json:"id"`
}

// 删除记录示例.
type DeleteReq struct {
	g.Meta `path:"/demo/:id" method:"delete" tags:"DemoService" summary:"Delete a demo record"`
	ID     uint `p:"id" in:"path" v:"required|integer|min:1"`
}

type DeleteRes struct {
}

// 更新记录示例.
type UpdateReq struct {
	g.Meta `path:"/demo/:id" method:"put" tags:"DemoService" summary:"Update a demo record"`
	ID     uint   `p:"id" in:"path" v:"integer|min:1"`
	Fielda string `p:"fileda" v:"required|passport|length:4,30"`
	Fieldb string `p:"filedb" v:"required|length:10,30"`
}

type UpdateRes struct {
}
