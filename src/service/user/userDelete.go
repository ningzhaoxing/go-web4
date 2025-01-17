package user

import (
	"userManageSystem-blog/src/dao/db/userDb"
	"userManageSystem-blog/src/model/params/sql"
	"userManageSystem-blog/src/model/user"
	"userManageSystem-blog/src/pkg/globals"
)

type UserDelete struct {
	err    error
	ut     user.User // 删除的目标用户
	uo     user.User // 操作用户
	appCtx *globals.AppCtx
}

func NewUserDelete(ut, uo user.User, appCtx *globals.AppCtx) *UserDelete {
	return &UserDelete{
		err:    nil,
		ut:     ut,
		uo:     uo,
		appCtx: appCtx,
	}
}

func (u *UserDelete) Delete() error {
	u.UserIsExist().InterceptOutPermission().DeleteUser()
	return u.err
}

// UserIsExist 验证用户是否存在
func (u *UserDelete) UserIsExist() *UserDelete {
	if u.err != nil {
		return u
	}
	ug := NewUserRegister(&u.ut, u.appCtx)
	_, u.err = ug.UserIsExist()
	return u
}

// InterceptOutPermission 验证操作权限(非管理员或本人操作)
func (u *UserDelete) InterceptOutPermission() *UserDelete {
	ud := NewUserEdit(u.ut, u.uo, u.appCtx)
	// 管理员或本人操作满足其一即可
	var b1, b2 bool
	b1, u.err = ud.InterceptNotManagerOpera()
	b2, u.err = ud.InterceptNotOwnOpera()
	if !b1 && !b2 {
		return u
	}
	u.err = nil
	return u
}

// DeleteUser 删除用户
func (u *UserDelete) DeleteUser() *UserDelete {
	if u.err != nil {
		return u
	}
	u.err = userDb.DeleteUser(sql.UserSqlParam{
		Db:   u.appCtx.GetDb(),
		User: u.ut,
	})
	return u
}
