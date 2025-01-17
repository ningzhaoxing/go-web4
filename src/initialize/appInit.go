package initialize

import (
	"userManageSystem-blog/src/pkg/globals"
)

func AppInit() {
	var err error
	globals.C, err = InitConfig()
	if err != nil {
		return
	}

	globals.Db = InitDb(globals.C)

	//globals.Rdb = InitRedis(globals.C)

	globals.EventBus = InitEventBus()
}
