// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure of table users for DAO operations like Where/Data.
type Users struct {
	g.Meta    `orm:"table:users, do:true"`
	Id        interface{} //
	Name      interface{} //
	Email     interface{} //
	Password  interface{} //
	Phone     interface{} //
	Role      interface{} // 用户权限角色
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
