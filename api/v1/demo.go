package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gf2-demo/internal/model"
)

/*
API文档: https://goframe.org/pages/viewpage.action?pageId=47703679
路由规则: https://goframe.org/pages/viewpage.action?pageId=1114257
请求输入: https://goframe.org/pages/viewpage.action?pageId=1114483
字段校验: https://goframe.org/pages/viewpage.action?pageId=1114367
*/

// 查询记录示例
type DemoGetReq struct {
	g.Meta `path:"/demo/:fielda" method:"get" tags:"DemoService" summary:"Get demo info by fielda"`
	Fielda string `p:"fielda" in:"path" v:"required|length:4,30"`
}

type DemoGetRes struct {
	ID        uint        `json:"id"`
	Fielda    string      `json:"fielda"`
	CreatedAt *gtime.Time `json:"created_at"`
	UpdatedAt *gtime.Time `json:"updated_at"`
}

// 创建记录示例
type DemoCreateReq struct {
	g.Meta `path:"/demo" method:"post" tags:"DemoService" summary:"Create a demo record"`
	Fielda string `p:"fileda" v:"required|passport|length:4,30"`
	Fieldb string `p:"filedb" v:"required|length:10,30"`
}

type DemoCreateRes struct {
	ID uint `json:"id"`
}

// 删除记录示例
type DemoDeleteReq struct {
	g.Meta `path:"/demo/:id" method:"delete" tags:"DemoService" summary:"Delete a demo record"`
	ID     uint `p:"id" in:"path" v:"required|integer|min:1"`
}

type DemoDeleteRes struct {
}

// 更新记录示例
type DemoUpdateReq struct {
	g.Meta `path:"/demo/:id" method:"put" tags:"DemoService" summary:"Update a demo record"`
	ID     uint   `p:"id" in:"path" v:"integer|min:1"`
	Fielda string `p:"fileda" v:"required|passport|length:4,30"`
	Fieldb string `p:"filedb" v:"required|length:10,30"`
}

type DemoUpdateRes struct {
}

// 查询记录列表示例
type DemoListReq struct {
	g.Meta `path:"/demo" method:"get" tags:"DemoService" summary:"Get demo records list"`
	CommonPaginationReq
}

type DemoListRes struct {
	CommonPaginationRes

	List []model.DemoListOutputItem `json:"list"`
}
