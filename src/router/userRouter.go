package router

import (
	"github.com/gorilla/mux"
	"userManageSystem-blog/src/controller/middleWare"
	"userManageSystem-blog/src/controller/user"
)

func userRouter(r *mux.Router) {

	r.HandleFunc("/user/login", user.UserLoginController)
	r.HandleFunc("/user/register", user.UserRegisterController)
	r.HandleFunc("/user/getOwnInfo", middleWare.AuthMiddleware(user.UserOwnInfoController))
	r.HandleFunc("/user/delete", middleWare.AuthMiddleware(user.UserDeleteController))
	r.HandleFunc("/user/edit", middleWare.AuthMiddleware(user.UserEditController))
	r.HandleFunc("/user/list", middleWare.AuthMiddleware(user.UserListController))
	r.HandleFunc("/user/info", middleWare.AuthMiddleware(user.UserInfoController))
	r.HandleFunc("/user/add", middleWare.AuthMiddleware(user.AddUserController))

	// restful风格
	// /users  "GET"
	// /users/1	"GET"
	// /users   "POST"
	// /users/1 "PUT"
	// /users/1	"DELETE"
	// /articles/1 "GET"
}
