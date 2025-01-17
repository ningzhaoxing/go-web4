package user

import (
	"userManageSystem-blog/src/dao/db/userDb"
	"userManageSystem-blog/src/model/params/sql"
	"userManageSystem-blog/src/model/user"
	"userManageSystem-blog/src/pkg/globals"
)

type UserInfo struct {
	err    error
	appCtx *globals.AppCtx
	user   user.User
}

func NewUserInfo(appCtx *globals.AppCtx, user user.User) *UserInfo {
	return &UserInfo{
		appCtx: appCtx,
		user:   user,
	}
}

func (u *UserInfo) GetUserInfo() (*user.User, error) {
	us, err := userDb.QueryUserInfoById(sql.UserSqlParam{
		Db:   u.appCtx.GetDb(),
		User: u.user,
	})
	if err != nil {
		return nil, err
	}
	return us, nil
}
