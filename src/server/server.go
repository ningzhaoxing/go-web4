package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"userManageSystem-blog/src/initialize"
	"userManageSystem-blog/src/pkg/globals"
	"userManageSystem-blog/src/router"
)

func Run() {
	initialize.AppInit()
	defer globals.Db.Close()

	r := mux.NewRouter()

	router.RunRouters(r)

	fmt.Println("启动成功...")
	err := http.ListenAndServe(fmt.Sprint(globals.C.App.Host, ":", globals.C.App.Port), r)

	if err != nil {
		return
	}
}
