package user

import (
	"userManageSystem-blog/src/dao/db/userDb"
	"userManageSystem-blog/src/model/params/service"
	"userManageSystem-blog/src/model/params/sql"
	"userManageSystem-blog/src/model/user"
	"userManageSystem-blog/src/pkg/globals"
)

// UserList 获取已注册用户列表
type UserList struct {
	err    error
	appCtx *globals.AppCtx
	uo     user.User
}

func NewUserList(appCtx *globals.AppCtx, uo user.User) *UserList {
	return &UserList{
		err:    nil,
		appCtx: appCtx,
		uo:     uo,
	}
}

func (u *UserList) UserList(p service.PageQueryParam) ([]user.User, error) {
	var users []user.User
	users = u.GetUserListByPage(p)

	return users, u.err
}

// GetUserListByPage 分页获取用户列表
func (u *UserList) GetUserListByPage(p service.PageQueryParam) []user.User {
	if u.err != nil {
		return nil
	}
	users := make([]user.User, 0)
	users, u.err = userDb.QueryAllUserByPage(sql.UserSqlParam{
		Db:             u.appCtx.GetDb(),
		User:           user.User{},
		PageQueryParam: p,
	})
	return users
}

// GetUserNum 获取所有已注册用户的数量
func (u *UserList) GetUserNum() int {
	var num int
	num, u.err = userDb.QueryUserNum(sql.UserSqlParam{
		Db: u.appCtx.GetDb(),
	})
	return num
}
