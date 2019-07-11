package router

import (
	"gonative/app/controller"
	"net/http"
)

/**
 * 初始化静态资源目录
 **/
func InitStaticDir(mux *http.ServeMux) {
	files := http.FileServer(http.Dir("app/views"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
}

/**
 * 初始化路由
 **/
func InitRoute() {

	http.Handle("/index", &controller.IndexController{})

}

/**
 * 加载函数
 **/
func InitFuncRoute() {
	http.HandleFunc("/index/login", controller.Login)
	http.HandleFunc("/index/json", controller.JsonReturn)

	http.HandleFunc("/basic/index", controller.Study)
}
