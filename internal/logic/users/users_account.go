package users

import (
	"context"
	"fmt"
	"framehub/internal/dao"
	"framehub/internal/model/entity"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
)

var JwtConfig struct {
	SecretKey string
	Issuer    string
	Expire    time.Duration
}

func InitJwtConfig(ctx context.Context) {
	JwtConfig.SecretKey = g.Cfg().MustGet(ctx, "jwt.secretKey").String()
	JwtConfig.Issuer = g.Cfg().MustGet(ctx, "jwt.issuer").String()
	expireHour := g.Cfg().MustGet(ctx, "jwt.expire").Int()
	JwtConfig.Expire = time.Duration(expireHour) * time.Hour
}

type jwtClaims struct {
	Id   uint
	Name string
	Role string
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
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    JwtConfig.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(JwtConfig.Expire)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	fmt.Println("JWT Secret Key:", JwtConfig.SecretKey)
	return token.SignedString([]byte(JwtConfig.SecretKey))
}

func (u *Users) Info(ctx context.Context) (user *entity.Users, err error) {
	// 从上下文中获取用户信息
	tokenString := g.RequestFromCtx(ctx).Request.Header.Get("Authorization")

	// 如果有 "Bearer " 前缀，移除它
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	tokenClaims, _ := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(JwtConfig.SecretKey), nil
	})

	if claims, ok := tokenClaims.Claims.(*jwtClaims); ok && tokenClaims.Valid {
		err = dao.Users.Ctx(ctx).Where("id", claims.Id).Scan(&user)
	}
	return
}
