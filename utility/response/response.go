package response

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// ShowType
const (
	SHOW_TYPE_SLIENT       = 0
	SHOW_TYPE_TIP          = 1
	SHOW_TYPE_WARNING      = 2
	SHOW_TYPE_NOTIFICATION = 3
	SHOW_TYPE_PAGE         = 4
)

// 数据结构
type JsonRes struct {
	Success      bool        `json:"success"`      // 是否成功
	Data         interface{} `json:"data"`         // 返回数据
	ErrorCode    string      `json:"errorCode"`    // 业务错误码
	ErrorMessage string      `json:"errorMessage"` // 错误提示
	ShowType     int         `json:"showType"`     // 错误类型 0=沉默 1=提示 2=警告 4=通知 5=错误页面
}

// Json 返回标准JSON数据。
func Json(r *ghttp.Request, success bool, showType int, code string, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	} else {
		responseData = g.Map{}
	}

	r.Response.WriteJson(JsonRes{
		Success:      success,
		Data:         responseData,
		ErrorCode:    code,
		ErrorMessage: message,
		ShowType:     showType,
	})
}

// JsonSuccess 返回标准的成功JSON数据。
func JsonSuccess(r *ghttp.Request, data ...interface{}) {
	Json(r, true, SHOW_TYPE_SLIENT, "0", "ok", data...)
}

// JsonSuccessExit 返回标准JSON数据并退出当前HTTP执行函数。
func JsonSuccessExit(r *ghttp.Request, data ...interface{}) {
	JsonSuccess(r, data...)
	r.Exit()
}

// JsonError 返回标准的错误JSON数据
func JsonError(r *ghttp.Request, code string, message string, data ...interface{}) {
	Json(r, false, SHOW_TYPE_WARNING, code, message, data...)
}

// JsonErrorExit 返回标准的错误JSON数据并退出当前HTTP执行函数。
func JsonErrorExit(r *ghttp.Request, code string, message string, data ...interface{}) {
	JsonError(r, code, message, data...)
	r.Exit()
}
