package users

import (
	"context"

	v1 "framehub/api/users/v1"
	"framehub/internal/logic/users"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	err = c.users.Register(ctx, users.RegisterInput{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
	})
	return nil, err
}
