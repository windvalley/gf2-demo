package model

import "github.com/gogf/gf/v2/os/gtime"

type DemoCreateInput struct {
	Fielda string
	Fieldb string
}

type DemoCreateOutput struct {
	ID uint
}

type DemoUpdateInput struct {
	ID     uint
	Fielda string
	Fieldb string
}

type DemoUpdateOutput struct {
}

type DemoListInput struct {
	PageNum  int
	PageSize int
}

type DemoListOutput struct {
	PageNum  int
	PageSize int
	Total    int
	List     []DemoListOutputItem
}

// DemoListOutputItem NOTE: 此处为了不返回Fieldb字段, 所以重新定义返回结构体, 否则可以直接使用enttity.Demo.
type DemoListOutputItem struct {
	ID        uint        `json:"id"`
	Fielda    string      `json:"fielda"`
	CreatedAt *gtime.Time `json:"created_at"`
	UpdatedAt *gtime.Time `json:"updated_at"`
}
