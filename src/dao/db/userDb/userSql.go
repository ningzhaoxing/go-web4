package userDb

import (
	user2 "userManageSystem-blog/src/model/params/sql"
	"userManageSystem-blog/src/model/user"
)

// QueryUserEmailByEmail 根据email查询指定用户email(用户判断用户是否存在)
func QueryUserEmailByEmail(p user2.UserSqlParam) (err error) {
	var user user.User
	sqlStr := "select email from users where email = ?"
	if err := p.Db.QueryRow(sqlStr, p.User.Email).Scan(&user.Email); err != nil {
		return err
	}
	return nil
}

// InsertIntoUsers 插入用户信息
func InsertIntoUsers(p user2.UserSqlParam) error {
	sqlStr := "insert into users(id,email,password,name,head_photo,gender) value (?,?,?,?,?,?)"
	_, err := p.Db.Exec(sqlStr, p.User.Id, p.User.Email, p.User.Password, p.User.Name, p.User.HeadPhoto, p.User.Gender)

	if err != nil {
		return err
	}
	return nil
}

// QueryUserPasswordByEmail 查询指定用户password
func QueryUserPasswordByEmail(p user2.UserSqlParam) (psd string, err error) {
	sqlStr := "select password from users where email = ?"
	if err := p.Db.QueryRow(sqlStr, p.User.Email).Scan(&psd); err != nil {
		return "", err
	}
	return psd, nil
}

// QueryUserByEmail 根据email查询用户所有信息
func QueryUserByEmail(p user2.UserSqlParam) (*user.User, error) {
	user := user.User{}
	sqlStr := "select id, email, password, name, gender, head_photo, permission_level from users where email = ?"
	if err := p.Db.QueryRow(sqlStr, p.User.Email).Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Gender, &user.HeadPhoto, &user.PermissionLevel); err != nil {
		return nil, err
	}
	return &user, nil
}

// QueryAllUserByPage 分页查询用户信息
func QueryAllUserByPage(p user2.UserSqlParam) ([]user.User, error) {
	users := make([]user.User, 0)
	sqlStr := "select id, email, password, name, gender, head_photo, permission_level from users where del = ? limit ?,? "
	rows, err := p.Db.Query(sqlStr, false, (p.PageQueryParam.Page-1)*p.PageQueryParam.Limit, p.PageQueryParam.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u user.User
		err := rows.Scan(&u.Id, &u.Email, &u.Password, &u.Name, &u.Gender, &u.HeadPhoto, &u.PermissionLevel)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// DeleteUser 删除用户
func DeleteUser(p user2.UserSqlParam) error {
	sqlStr := "update users set del	= true where id = ?"
	_, err := p.Db.Exec(sqlStr, p.User.Id)
	if err != nil {
		return err
	}
	return err
}

// UpdateUser 修改用户信息
func UpdateUser(p user2.UserSqlParam) error {
	sqlStr := "update users set password = ?, name = ?, head_photo = ?, gender = ? where id = ?"
	_, err := p.Db.Exec(sqlStr, p.User.Password, p.User.Name, p.User.HeadPhoto, p.User.Gender, p.User.Id)
	if err != nil {
		return err
	}
	return nil
}

// QueryUserNum 获取当前已注册用户的数量
func QueryUserNum(p user2.UserSqlParam) (int, error) {
	var count int
	sqlStr := "select count(*) from users where del = ?"
	if err := p.Db.QueryRow(sqlStr, false).Scan(&count); err != nil {
		return -1, err
	}
	return count, nil
}

// QueryUserInfoById 查看用户信息通过id
func QueryUserInfoById(p user2.UserSqlParam) (*user.User, error) {
	var u user.User
	sqlStr := "select id, email, password, name, gender, head_photo, permission_level from users where id = ?"
	if err := p.Db.QueryRow(sqlStr, p.User.Id).Scan(&u.Id, &u.Email, &u.Password, &u.Name, &u.Gender, &u.HeadPhoto, &u.PermissionLevel); err != nil {
		return nil, err
	}
	return &u, nil
}

// UpdateUserHeadPhoto 修改用户头像
func UpdateUserHeadPhoto(p user2.UserSqlParam) error {
	sqlStr := "update users set head_photo = ? where id = ?"
	if _, err := p.Db.Exec(sqlStr, p.User.HeadPhoto, p.User.Id); err != nil {
		return err
	}
	return nil
}
