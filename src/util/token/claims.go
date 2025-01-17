package token

import (
	"github.com/dgrijalva/jwt-go"
	"userManageSystem-blog/src/model/params/util"
	"userManageSystem-blog/src/model/user"
)

type CustomClaims struct {
	User user.User
	jwt.StandardClaims
}

func NewCustomClaims(c util.CustomClaimsParam) *CustomClaims {
	c.User.Password = ""
	return &CustomClaims{
		User:           c.User,
		StandardClaims: c.Jwt,
	}
}
