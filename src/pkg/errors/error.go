package errors

import "userManageSystem-blog/src/pkg/globals"

var (
	ErrUserHasExist        = NewAppError(globals.CodeHttpBad, "用户已存在")
	ErrInternalServer      = NewAppError(globals.CodeSercerError, "服务端内部错误")
	ErrOriginNotAllowed    = NewAppError(globals.CodeFailed, "请求源被拒绝")
	ErrHttpHeader          = NewAppError(globals.CodeHttpBad, "请求头错误")
	ErrDataForm            = NewAppError(globals.CodeHttpBad, "数据格式错误")
	ErrAccountOrPsd        = NewAppError(globals.CodeHttpBad, "邮箱或密码错误")
	ErrToken               = NewAppError(globals.CodeFailed, "token不合法")
	ErrNotFoundUserInToken = NewAppError(globals.CodeSercerError, "无法从token中获取user")
	ErrPermissionsOut      = NewAppError(globals.CodeUnauthorized, "权限越界")
)
