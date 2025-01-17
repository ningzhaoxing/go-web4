package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RunRouters(r *mux.Router) {

	r.PathPrefix("/plugins/").Handler(http.StripPrefix("/plugins/", http.FileServer(http.Dir("views/static/plugins"))))
	r.PathPrefix("/pages/").Handler(http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("views/static/css"))))
	r.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("views/static/img"))))
	r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))
	// 注册用户路由
	userRouter(r)

	// 注册文件路由
	fileRouter(r)
}
