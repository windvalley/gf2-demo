package codes

import "github.com/gogf/gf/v2/errors/gcode"

type BizCode struct {
	code    int
	message string
	detail  BizCodeDetail
}

type BizCodeDetail struct {
	Code     string
	HTTPCode int
}

func (c BizCode) BizDetail() BizCodeDetail {
	return c.detail
}

func (c BizCode) Code() int {
	return c.code
}

func (c BizCode) Message() string {
	return c.message
}

func (c BizCode) Detail() interface{} {
	return c.detail
}

func New(httpCode int, code, message string) gcode.Code {
	return BizCode{
		code:    0,
		message: message,
		detail: BizCodeDetail{
			Code:     code,
			HTTPCode: httpCode,
		},
	}
}
