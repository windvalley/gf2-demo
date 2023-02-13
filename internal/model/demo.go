package model

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
