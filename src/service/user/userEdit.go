package user

import (
	"userManageSystem-blog/src/dao/db/userDb"
	"userManageSystem-blog/src/model/params/sql"
	"userManageSystem-blog/src/model/user"
	"userManageSystem-blog/src/pkg/errors"
	"userManageSystem-blog/src/pkg/globals"
)

type UserEdit struct {
	err    error
	ut     user.User
	uo     user.User
	appCtx *globals.AppCtx
}

func NewUserEdit(ut, uo user.User, appCtx *globals.AppCtx) *UserEdit {
	return &UserEdit{
		err:    nil,
		ut:     ut,
		uo:     uo,
		appCtx: appCtx,
	}
}

func (u *UserEdit) Edit() error {
	u.haveCheck().UpdateUserInfor()
	return u.err
}

func (u *UserEdit) haveCheck() *UserEdit {
	u.err = NewUserDelete(u.ut, u.uo, u.appCtx).InterceptOutPermission().err
	if u.err != nil {
		return u
	}
	_, u.err = u.CheckRequiredFieldIsNull()
	return u
}

// InterceptNotManagerOpera 拦截非管理员操作
func (u *UserEdit) InterceptNotManagerOpera() (bool, error) {
	// 不是管理员
	if u.uo.PermissionLevel != 1 {
		return false, errors.ErrPermissionsOut
	}
	return true, nil
}

// InterceptNotOwnOpera 拦截非本人操作
func (u *UserEdit) InterceptNotOwnOpera() (bool, error) {
	if u.uo.Email != u.ut.Email {
		return false, errors.ErrPermissionsOut
	}
	return true, nil
}

// UpdateUserInfor 修改信息
func (u *UserEdit) UpdateUserInfor() *UserEdit {
	if u.err != nil {
		return u
	}

	u.err = userDb.UpdateUser(sql.UserSqlParam{
		Db:   u.appCtx.GetDb(),
		User: u.ut,
	})
	return u
}

// CheckRequiredFieldIsNull 检查必填字段
func (u *UserEdit) CheckRequiredFieldIsNull() (bool, error) {
	return NewUserAdd(u.ut, u.uo, u.appCtx).CheckRequiredFieldIsNull()
}
