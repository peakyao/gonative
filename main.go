package main

import (
	"gonative/libs/slog"
	"gonative/router"
	"net/http"
)

var (
	mux *http.ServeMux
)

func init() {
	slog.InfoLog("start init")

	// router.InitStaticDir(mux)

}

func main() {
	slog.InfoLog("start ========= main")
	// slog.FuncLog(log)
	server := http.Server{
		Addr: "127.0.0.1:8080",
		// 	// Handler: &handler,
		// 	Handler: mux,
	}
	router.InitRoute()

	router.InitFuncRoute()

	server.ListenAndServe()
}

func log(w http.ResponseWriter, r *http.Request) {
	slog.InfoLog("====log=====")
}
