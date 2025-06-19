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

// Register 用户注册逻辑
func (u *Users) Register(ctx context.Context, in RegisterInput) error {
	if err := u.checkUser(ctx, in.Name, in.Email); err != nil {
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

// checkUser 检查用户名和邮箱是否已存在
func (u *Users) checkUser(ctx context.Context, name, email string) error {
	count, err := dao.Users.Ctx(ctx).Where("name", name).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("用户已存在")
	}
	count, err = dao.Users.Ctx(ctx).Where("email", email).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("邮箱已存在")
	}
	return nil
}
