package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type InfoReq struct {
	g.Meta `path:"account/info" method:"get" sm:"获取信息" tags:"用户"`
}

type InfoRes struct {
	Name      string      `json:"name" dc:"用户名"`
	Email     string      `json:"email" dc:"邮箱"`
	Role      string      `json:"role" dc:"角色"`
	CreatedAt *gtime.Time `json:"created" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"update" dc:"更新时间"`
}
