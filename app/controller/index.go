package controller

import (
	"gonative/libs/slog"
	"net/http"
	"text/template"
)

type IndexController struct {
}

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
