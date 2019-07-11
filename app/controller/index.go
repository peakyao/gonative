package controller

import (
	"encoding/json"
	"gonative/libs/slog"
	"net/http"
	"text/template"
)

type IndexController struct {
}

/**
 * 加载Handle路由
 **/
func (this *IndexController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slog.InfoLog("Index-index")

	len := r.ContentLength
	slog.InfoLog("contentLength", len)

	userAgent := r.UserAgent
	slog.InfoLog("userAgent", userAgent)

	referer := r.Referer
	slog.InfoLog("referer", referer)

	header := r.Header
	slog.InfoLog("header", header)

	requestURI := r.RequestURI
	slog.InfoLog("requestURI", requestURI)

	remoteAddr := r.RemoteAddr
	slog.InfoLog("remoteAddr", remoteAddr)

	// cookie := r.Cookies()
	// slog.InfoLog("cookie str ", &cookie.String())

	t, ok := template.ParseFiles("index.html")
	if ok != nil {
		slog.InfoLog("ERRORR======")
	}
	t.Execute(w, "hello=====================ServeHTTP")

	// w.Write([]byte("IndexaaaaaaaaIndex"))
}

/**
 * 登陆方法
 **/
func Login(w http.ResponseWriter, r *http.Request) {
	slog.InfoLog("Index-Login")
	w.Write([]byte("Index/Login"))
}

func JsonReturn(w http.ResponseWriter, r *http.Request) {
	// slog.InfoLog("Index-JsonReturn")
	var data string
	data = "ssssssssssss"
	// data["a"] = 1
	// data["b"] = 2

	// res := encode.ReturnResult{
	// 	Data:   data,
	// 	Errno:  0,
	// 	ErrMsg: "success",
	// }
	// result, _ := encode.Jsonencode(data)
	result, _ := json.Marshal(data)
	// w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
