package util

import (
	"github.com/dgrijalva/jwt-go"
	"userManageSystem-blog/src/model/user"
)

type CustomClaimsParam struct {
	User user.User          `json:"sql"`
	Jwt  jwt.StandardClaims `json:"jwt"`
}
