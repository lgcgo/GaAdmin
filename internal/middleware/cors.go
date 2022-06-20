package middleware

import "github.com/gogf/gf/v2/net/ghttp"

// 跨域请求中间件
func CORS(r *ghttp.Request) {
	var (
		options ghttp.CORSOptions
	)

	options = r.Response.DefaultCORSOptions()
	// 这里自定义跨域请求选项
	// ...

	r.Response.CORS(options)

	r.Middleware.Next()
}
