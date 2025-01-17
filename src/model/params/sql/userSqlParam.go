package sql

import (
	"database/sql"
	"userManageSystem-blog/src/model/params/service"
	"userManageSystem-blog/src/model/user"
)

type UserSqlParam struct {
	Db             *sql.DB
	User           user.User
	PageQueryParam service.PageQueryParam
}
