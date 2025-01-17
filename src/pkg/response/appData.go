package response

import (
	"userManageSystem-blog/src/pkg/globals"
)

type Appdata struct {
	Code globals.AppCode `json:"code"`
	Msg  string          `json:"msg"`
	Data interface{}     `json:"data"`
	Temp string          `json:"temp"`
}
