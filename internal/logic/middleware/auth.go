package middleware

import (
	"framehub/internal/consts"
	"net/http"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware 返回一个可配置的中间件，要求最低权限为 minRole
func AuthMiddleware(minRole consts.UserRole) ghttp.HandlerFunc {
	return func(r *ghttp.Request) {
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

		// 从 token 的 claims 中获取用户角色
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			r.Response.WriteStatus(http.StatusInternalServerError)
			r.Exit()
		}

		// 获取用户角色并转换为 consts.UserRole 类型
		roleStr, ok := claims["role"].(string)
		if !ok {
			r.Response.WriteStatus(http.StatusForbidden)
			r.Exit()
		}
		var userRole consts.UserRole
		switch roleStr {
		case "guest":
			userRole = consts.Guest
		case "user":
			userRole = consts.User
		case "admin":
			userRole = consts.Admin
		case "operator":
			userRole = consts.Operator
		default:
			r.Response.WriteStatus(http.StatusForbidden)
			r.Exit()
		}

		// 检查权限
		if userRole < minRole {
			r.Response.WriteStatus(http.StatusForbidden)
			r.Exit()
		}

		r.Middleware.Next()
	}
}
