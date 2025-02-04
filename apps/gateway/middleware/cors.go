package middleware

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/cors"
)

// Cors returns a CORS middleware handler
func Cors() app.HandlerFunc {
	return cors.New(cors.Config{
		// 允许的域名，这里设置具体的域名
		AllowOrigins: []string{
			"http://localhost:5173", // 前端开发服务器地址
			"http://127.0.0.1:5173",
		},
		// 允许的请求方法
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD", "PATCH",
		},
		// 允许的请求头
		AllowHeaders: []string{
			"Origin", "Content-Type", "Accept", "Authorization",
			"X-Requested-With", "X-CSRF-Token", "Token",
			"Access-Control-Allow-Origin", "Access-Control-Allow-Headers",
			"Access-Control-Allow-Methods", "Access-Control-Allow-Credentials",
		},
		// 是否允许携带 cookie
		AllowCredentials: true,
		// 预检请求的有效期，单位秒
		MaxAge: 86400,
		// 允许暴露的响应头
		ExposeHeaders: []string{
			"Content-Length",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Access-Control-Allow-Methods",
			"Access-Control-Allow-Credentials",
		},
	})
}
