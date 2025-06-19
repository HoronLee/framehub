package users

import (
	"context"
	v1 "framehub/api/users/v1"
)

// Login 用户登录控制器
func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	token, err := c.users.Login(ctx, req.Name, req.Password)
	if err != nil {
		return
	}
	return &v1.LoginRes{Token: token}, nil
}
