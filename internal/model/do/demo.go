// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Demo is the golang structure of table demo for DAO operations like Where/Data.
type Demo struct {
	g.Meta    `orm:"table:demo, do:true"`
	Id        interface{} // ID
	Fielda    interface{} // Field demo
	Fieldb    interface{} // Private field demo
	CreatedAt *gtime.Time // Created Time
	UpdatedAt *gtime.Time // Updated Time
}
