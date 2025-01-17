package user

import (
	"fmt"
	"userManageSystem-blog/src/model/user"
	"userManageSystem-blog/src/pkg/errors"
	"userManageSystem-blog/src/pkg/globals"
)

type UserAdd struct {
	err    error
	ut     user.User // 添加的目标用户
	uo     user.User // 操作用户
	appCtx *globals.AppCtx
}

func NewUserAdd(ut user.User, uo user.User, appCtx *globals.AppCtx) *UserAdd {
	return &UserAdd{
		ut:     ut,
		uo:     uo,
		appCtx: appCtx,
	}
}

func (u *UserAdd) AddUser() error {
	u.havaCheck()
	if u.err != nil {
		return u.err
	}
	return NewUserRegister(&u.ut, u.appCtx).GenerateUuid().SaveInDb().err
}

// 用户添加前的数据检查
func (u *UserAdd) havaCheck() *UserAdd {
	_, u.err = u.checkManagerOperateOrNot()
	if u.err != nil {
		return u
	}
	_, u.err = u.CheckRequiredFieldIsNull()
	if u.err != nil {
		return u
	}
	_, u.err = u.checkEmailFormat()
	if u.err != nil {
		return u
	}
	_, u.err = u.userIsExist()
	if u.err != nil {
		return u
	}
	return nil
}

// CheckManagerOperateOrNot 检查是否为管理员操作
func (u *UserAdd) checkManagerOperateOrNot() (bool, error) {
	if u.uo.PermissionLevel != 1 {
		return false, errors.ErrPermissionsOut
	}
	return true, nil
}

// CheckRequiredFieldIsNull 检查必填字段是否为空(邮箱、密码、姓名)
func (u *UserAdd) CheckRequiredFieldIsNull() (bool, error) {
	if u.ut.Email == "" || u.ut.Password == "" || u.ut.Name == "" {
		fmt.Println(u.ut)
		return true, errors.ErrDataForm
	}
	return false, nil
}

// 检查是否存在该用户
func (u *UserAdd) userIsExist() (bool, error) {
	return NewUserRegister(&u.ut, u.appCtx).UserIsExist()
}

// 检查邮箱密码格式
func (u *UserAdd) checkEmailFormat() (bool, error) {
	return NewUserRegister(&u.ut, u.appCtx).CheckEmailAndPasswordFormat()
}
