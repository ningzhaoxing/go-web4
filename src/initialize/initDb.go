package initialize

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"userManageSystem-blog/src/pkg/globals"
)

func InitDb(c *globals.Config) *sql.DB {
	// DSN:Data Source Name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", c.DB.Config.User, c.DB.Config.Password, c.DB.Config.Host, c.DB.Config.Port, c.DB.Config.Db)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprint("Mysql connect err:", err.Error()))
		return nil
	}

	err = db.Ping()
	if err != nil {
		panic(fmt.Sprint("Mysql connect err:", err.Error()))
		return nil
	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	db.SetMaxIdleConns(c.DB.Config.MaxIdle)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.SetMaxOpenConns(c.DB.Config.MaxCon)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	db.SetConnMaxLifetime(time.Hour)

	return db
}
