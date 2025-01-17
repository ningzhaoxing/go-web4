package user

import (
	"database/sql"
	"errors"
	"userManageSystem-blog/src/dao/db/userDb"
	sql2 "userManageSystem-blog/src/model/params/sql"
	"userManageSystem-blog/src/model/user"
	errors2 "userManageSystem-blog/src/pkg/errors"
	"userManageSystem-blog/src/pkg/globals"
	"userManageSystem-blog/src/util/uuid"
)

type UserRegister struct {
	err    error
	user   *user.User
	appCtx *globals.AppCtx
}

func NewUserRegister(user *user.User, appCtx *globals.AppCtx) *UserRegister {
	return &UserRegister{
		err:    nil,
		user:   user,
		appCtx: appCtx,
	}
}

func (ug *UserRegister) Register() error {
	var isExist, ok bool
	isExist, ug.err = ug.UserIsExist()
	if isExist {
		return errors2.ErrUserHasExist
	}
	ok, ug.err = ug.CheckEmailAndPasswordFormat()
	if !ok {
		return errors2.ErrDataForm
	}
	ug.GenerateUuid().SaveInDb()
	return nil
}

// GenerateUuid 生成uuid给user
func (ug *UserRegister) GenerateUuid() *UserRegister {
	if ug.err != nil {
		return ug
	}
	ug.user.Id, ug.err = uuid.GetUuid()
	return ug
}

// SaveInDb 将用户信息添加到数据库
func (ug *UserRegister) SaveInDb() *UserRegister {
	if ug.err != nil {
		return ug
	}

	ug.err = userDb.InsertIntoUsers(sql2.UserSqlParam{
		Db:   ug.appCtx.GetDb(),
		User: *ug.user,
	})
	return ug
}

// UserIsExist 判断用户是否存在
func (ug *UserRegister) UserIsExist() (bool, error) {
	err := userDb.QueryUserEmailByEmail(sql2.UserSqlParam{
		Db:   ug.appCtx.GetDb(),
		User: *ug.user,
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// CheckEmailAndPasswordFormat CheckEmailFormat 邮箱格式验证
func (ug *UserRegister) CheckEmailAndPasswordFormat() (bool, error) {
	err := NewUserLogin(ug.user.Email, ug.user.Password, ug.appCtx).CheckFormat().err
	if err != nil {
		return false, err
	}
	return true, nil
}
