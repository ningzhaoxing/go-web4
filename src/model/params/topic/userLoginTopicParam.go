package topic

import (
	"userManageSystem-blog/src/model/user"
	"userManageSystem-blog/src/pkg/response"
)

type UserLoginTopicParam struct {
	R    *response.Response
	User user.User
}
