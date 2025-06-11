package users

import (
	"context"
	"framehub/internal/consts"
	"framehub/internal/dao"
	"framehub/internal/model/entity"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/golang-jwt/jwt/v5"
)

type jwtClaims struct {
	Id   uint
	Name string
	jwt.RegisteredClaims
}

func (u *Users) Login(ctx context.Context, name, password string) (tokenString string, err error) {
	var user entity.Users
	err = dao.Users.Ctx(ctx).Where("name", name).Scan(&user)
	if err != nil {
		return "", gerror.New("查询用户失败")
	}

	if user.Id == 0 {
		return "", gerror.New("用户不存在")
	}

	// 将密码加密后与数据库中的密码进行比对
	if user.Password != u.encryptPassword(password) {
		return "", gerror.New("用户名或密码错误")
	}

	// 生成token
	uc := &jwtClaims{
		Id:   uint(user.Id),
		Name: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    consts.JwtIssuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	return token.SignedString([]byte(consts.JwtKey))
}
