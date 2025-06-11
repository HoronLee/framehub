// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	Id        uint64      `json:"id"        orm:"id"         description:""`       //
	Name      string      `json:"name"      orm:"name"       description:""`       //
	Email     string      `json:"email"     orm:"email"      description:""`       //
	Password  string      `json:"password"  orm:"password"   description:""`       //
	Phone     string      `json:"phone"     orm:"phone"      description:""`       //
	Role      string      `json:"role"      orm:"role"       description:"用户权限角色"` // 用户权限角色
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`       //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`       //
}
