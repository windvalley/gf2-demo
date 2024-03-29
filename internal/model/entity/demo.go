// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Demo is the golang structure for table demo.
type Demo struct {
	Id        uint        `json:"id"         description:"ID"`
	Fielda    string      `json:"fielda"     description:"Field demo"`
	Fieldb    string      `json:"fieldb"     description:"Private field demo"`
	CreatedAt *gtime.Time `json:"created_at" description:"Created Time"`
	UpdatedAt *gtime.Time `json:"updated_at" description:"Updated Time"`
}
