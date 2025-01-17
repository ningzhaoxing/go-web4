package user

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	service2 "userManageSystem-blog/src/model/params/service"
	"userManageSystem-blog/src/model/user"
	errors2 "userManageSystem-blog/src/pkg/errors"
	"userManageSystem-blog/src/pkg/globals"
	gx2 "userManageSystem-blog/src/pkg/gx"
	user2 "userManageSystem-blog/src/service/user"
	"userManageSystem-blog/src/util/pageQuery"
	"userManageSystem-blog/src/util/token"
)

//  验证
//	token  why,how，生成解析加密原理过程；对称加密和非对称加密
//  全局panic处理	做中间件

// 改进:
// 生命周期
// 安全：CSRF
// 登录注册限流

// UserLoginController 登录
func UserLoginController(w http.ResponseWriter, r *http.Request) {
	gx := gx2.NewDefaultGx(w, r)
	appCtx := globals.NewDefaultAppCtx()

	t := gx.GetTemplate().ParseFilesLogin()
	u := &user.User{}

	err := gx.GetBind().BindForm(u)
	if err != nil {
		fmt.Println("controller.userController.login.Register() err=", err)
		gx.GetResponse().Fail(err, t, "login")
		return
	}

	login := user2.NewUserLogin(u.Email, u.Password, appCtx)

	if err := login.Login(); err != nil {
		fmt.Println("controller.userController.login.Register() err=", err)
		gx.GetResponse().Fail(err, t, "login")
		return
	}

	tokenString, err := token.GetToken(*u)
	if err != nil {
		fmt.Println("controller.userController.login.GetToken() err=", err)
		gx.GetResponse().Fail(err, t, "login")
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "token", Value: tokenString, Path: "/", Expires: time.Now().Add(24 * time.Hour)})
	http.Redirect(w, r, "/user/list", http.StatusFound)
}

// UserRegisterController 注册
func UserRegisterController(w http.ResponseWriter, r *http.Request) {
	gx := gx2.NewDefaultGx(w, r)
	appCtx := globals.NewDefaultAppCtx()

	t := gx.GetTemplate().ParseFilesRegister()
	u := &user.User{}

	err := gx.GetBind().BindForm(u)
	if err != nil {
		fmt.Println("controller.userController.login.Register() err=", err)
		gx.GetResponse().Fail(err, t, "register")
		return
	}

	register := user2.NewUserRegister(u, appCtx)
	err = register.Register()
	if err != nil {
		fmt.Println("controller.userController.register err=", err)
		gx.GetResponse().Fail(err, t, "register")
		return
	}

	tokenString, err := token.GetToken(*u)
	if err != nil {
		fmt.Println("controller.userController.GetToken err=", err)
		gx.GetResponse().Fail(err, t, "register")
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "token", Value: tokenString, Path: "/", Expires: time.Now().Add(24 * time.Hour)})
	http.Redirect(w, r, "/user/list", http.StatusFound)
}

// UserListController 用户分页列表
func UserListController(w http.ResponseWriter, r *http.Request) {
	gx := gx2.NewDefaultGx(w, r)
	appCtx := globals.NewDefaultAppCtx()

	t := gx.GetTemplate().ParseFilesDefault()

	curPage := r.FormValue("pageNo")
	if curPage == "" {
		curPage = "1"
	}
	cur, err := strconv.Atoi(curPage)
	if err != nil {
		fmt.Println("controller.userController.UserList err=", err)
		gx.GetResponse().Fail(err, t, "list")
		return
	}

	uo, err := token.NewToken(r).GetUser(appCtx.GetDb())
	if err != nil {
		fmt.Println("controller.userController.GetUser err=", err)
		gx.GetResponse().Fail(err, t, "list")
		return
	}

	userListService := user2.NewUserList(appCtx, *uo)
	list, err := userListService.UserList(service2.PageQueryParam{
		Page:  cur,
		Limit: 7,
	})
	if err != nil {
		fmt.Println("controller.userController.UserList err=", err)
		gx.GetResponse().Fail(err, t, "list")
		return
	}
	totalNum := userListService.GetUserNum()

	page := pageQuery.NewPage(cur, 7, totalNum)

	gx.GetResponse().Success(map[string]any{
		"Content": map[string]any{
			"List": list,
			"Page": page,
		},
		"Header":  map[string]any{"User": uo},
		"Sidebar": map[string]any{"User": uo},
	}, t, "list")
}

// UserOwnInfoController 获取本人用户信息(直接通过解析token获取)
func UserOwnInfoController(w http.ResponseWriter, r *http.Request) {
	gx := gx2.NewDefaultGx(w, r)
	appCtx := globals.NewDefaultAppCtx()

	t := gx.GetTemplate().ParseSingleFile("views/pages/users/info.html")

	token := token.NewToken(r)
	u, err := token.GetUser(appCtx.GetDb())
	if err != nil {
		fmt.Println("controller.UserOwnInfoController.GetUser err=", err)
		gx.GetResponse().Fail(err, t, "own_info")
		return
	}

	if u == nil {
		fmt.Println("controller.UserOwnInfoController.GetUser err=", err)
		gx.GetResponse().Fail(errors2.ErrToken, t, "own_info")
		return
	}

	gx.GetResponse().Success(map[string]any{
		"OwnUser": u,
	}, t, "own_info")
}

// UserDeleteController 用户删除
func UserDeleteController(w http.ResponseWriter, r *http.Request) {
	gx := gx2.NewDefaultGx(w, r)
	appCtx := globals.NewDefaultAppCtx()
	ut := &user.User{}

	err := gx.GetBind().BindQuery(ut)
	if err != nil {
		fmt.Println("controller.userController.BindQuery err=", err)
		gx.GetResponse().HttpFail(err)
		return
	}

	uo, err := token.NewToken(r).GetUser(appCtx.GetDb())
	if err != nil {
		fmt.Println("controller.userController.GetUser err=", err)
		gx.GetResponse().HttpFail(err)
		return
	}

	err = user2.NewUserDelete(*ut, *uo, appCtx).Delete()
	if err != nil {
		fmt.Println("controller.userController.NewUserDelete err=", err)
		gx.GetResponse().HttpFail(err)
		return
	}

	gx.GetResponse().HttpSuccess("删除成功", struct{}{})
}

// UserEditController 用户信息编辑
func UserEditController(w http.ResponseWriter, r *http.Request) {
	gx := gx2.NewDefaultGx(w, r)
	appCtx := globals.NewDefaultAppCtx()
	ut := &user.User{}

	err := gx.GetBind().BindForm(ut)
	if err != nil {
		fmt.Println("controller.userController.GetUser err=", err)
		gx.GetResponse().HttpFail(err)
		return
	}
	uo, err := token.NewToken(r).GetUser(appCtx.GetDb())
	if err != nil {
		fmt.Println("controller.userController.GetUser err=", err)
		gx.GetResponse().HttpFail(err)
		return
	}

	err = user2.NewUserEdit(*ut, *uo, appCtx).Edit()
	if err != nil {
		fmt.Println("controller.userController.NewUserEdit err=", err)
		gx.GetResponse().HttpFail(err)
		return
	}

	gx.GetResponse().HttpSuccess("修改成功", struct{}{})
}

// UserInfoController 获取用户信息
func UserInfoController(w http.ResponseWriter, r *http.Request) {
	appCtx := globals.NewDefaultAppCtx()
	gx := gx2.NewDefaultGx(w, r)
	ut := user.User{}

	t := gx.GetTemplate().ParseFilesDefault()
	ownUser, err := token.NewToken(r).GetUser(appCtx.GetDb())
	if err != nil {
		fmt.Println("controller.UserInfoController.GetUserInfo err=", err)
		gx.GetResponse().Fail(err, t, "info")
		return
	}

	ut.Id = r.URL.Query().Get("id")
	u, err := user2.NewUserInfo(appCtx, ut).GetUserInfo()
	if err != nil {
		fmt.Println("controller.UserInfoController.GetUserInfo err=", err)
		gx.GetResponse().Fail(err, t, "info")
		return
	}

	gx.GetResponse().Success(map[string]any{
		"Content": map[string]any{"User": u},
		"Header":  map[string]any{"User": ownUser},
		"Sidebar": map[string]any{"User": ownUser},
	}, t, "info")
}

// AddUserController 添加用户
func AddUserController(w http.ResponseWriter, r *http.Request) {
	appCtx := globals.NewDefaultAppCtx()
	gx := gx2.NewDefaultGx(w, r)
	ut := &user.User{}

	err := gx.GetBind().BindForm(ut)
	if err != nil {
		fmt.Println("controller.userController.GetUser err=", err)
		gx.GetResponse().HttpFail(err)
		return
	}

	uo, err := token.NewToken(r).GetUser(appCtx.GetDb())
	if err != nil {
		fmt.Println("controller.userController.GetUser err=", err)
		gx.GetResponse().HttpFail(err)
		return
	}

	err = user2.NewUserAdd(*ut, *uo, appCtx).AddUser()
	if err != nil {
		fmt.Println("controller.userController.GetUser err=", err)
		gx.GetResponse().HttpFail(err)
		return
	}
	gx.GetResponse().HttpSuccess("修改成功", struct{}{})
}
