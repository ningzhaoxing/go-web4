package user

import (
	"regexp"
	"userManageSystem-blog/src/dao/db/userDb"
	"userManageSystem-blog/src/model/params/sql"
	"userManageSystem-blog/src/model/user"
	"userManageSystem-blog/src/pkg/errors"
	"userManageSystem-blog/src/pkg/globals"
)

// 判断邮箱和密码是否正确
// 是：登录成功	否：登录失败

type UserLogin struct {
	err    error
	user   user.User
	appCtx *globals.AppCtx
}

func NewUserLogin(email, psd string, appCtx *globals.AppCtx) *UserLogin {
	return &UserLogin{
		err:    nil,
		appCtx: appCtx,
		user: user.User{
			Email:    email,
			Password: psd,
		},
	}
}

func (u *UserLogin) Login() error {
	_, u.err = u.CheckFormat().CheckIdentity()
	if u.err != nil {
		return u.err
	}
	return nil
}

// CheckFormat 邮箱密码格式验证
func (u *UserLogin) CheckFormat() *UserLogin {
	if u.err != nil {
		return u
	}
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)

	if !(emailRegex.MatchString(u.user.Email) && len(u.user.Password) >= 6) {
		u.err = errors.ErrDataForm
	}
	return u
}

// CheckIdentity 检查邮箱和密码是否正确
func (u *UserLogin) CheckIdentity() (bool, error) {
	if u.err != nil {
		return false, u.err
	}

	var psd string
	psd, u.err = userDb.QueryUserPasswordByEmail(sql.UserSqlParam{
		Db:   u.appCtx.GetDb(),
		User: u.user,
	})
	if psd != u.user.Password {
		return false, errors.ErrAccountOrPsd
	}
	return true, nil
}
