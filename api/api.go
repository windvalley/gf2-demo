package api

type CommonPaginationReq struct {
	PageNum  int `json:"page_num" in:"query" v:"min:0#分页号码错误"  dc:"分页号码, 默认1"`
	PageSize int `json:"page_size" in:"query" v:"max:500#每页数量最大500条" dc:"分页数量, 最大500"`
}

type CommonPaginationRes struct {
	Total    int `json:"total" dc:"总数"`
	PageNum  int `json:"page_num,omitempty" dc:"分页号码"`
	PageSize int `json:"page_size,omitempty" dc:"分页数量"`
}
