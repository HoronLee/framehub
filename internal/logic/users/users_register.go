package users

import (
	"context"

	"framehub/internal/dao"
	"framehub/internal/model/do"

	"github.com/gogf/gf/v2/errors/gerror"
)

type RegisterInput struct {
	Name     string
	Password string
	Email    string
}

func (u *Users) Register(ctx context.Context, in RegisterInput) error {
	if err := u.checkUser(ctx, in.Name); err != nil {
		return err
	}
	_, err := dao.Users.Ctx(ctx).Data(do.Users{
		Name:     in.Name,
		Password: u.encryptPassword(in.Password),
		Email:    in.Email,
	}).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (u *Users) checkUser(ctx context.Context, name string) error {
	count, err := dao.Users.Ctx(ctx).Where("name", name).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("用户已存在")
	}
	return nil
}
