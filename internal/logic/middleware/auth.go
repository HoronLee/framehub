package middleware

import (
	"framehub/internal/consts"
	"net/http"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(r *ghttp.Request) {
	// 获取 Authorization 头部的值
	var authHeader = r.Header.Get("Authorization")

	// 去掉 "Bearer " 前缀
	var tokenString string
	if strings.HasPrefix(authHeader, "Bearer ") {
		tokenString = strings.TrimPrefix(authHeader, "Bearer ")
	} else {
		tokenString = authHeader
	}

	// 解析 JWT Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(consts.JwtKey), nil
	})
	if err != nil || !token.Valid {
		r.Response.WriteStatus(http.StatusForbidden)
		r.Exit()
	}
	r.Middleware.Next()
}
