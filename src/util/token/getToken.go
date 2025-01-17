package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"userManageSystem-blog/src/model/params/util"
	"userManageSystem-blog/src/model/user"
)

// CustomSecret 定义密钥
var CustomSecret = []byte("xlszxjm")

// GetToken 生成token
func GetToken(user user.User) (string, error) {
	tokenExpireDuration := time.Now().Add(7 * 24 * time.Hour)
	claims := NewCustomClaims(util.CustomClaimsParam{
		User: user, Jwt: jwt.StandardClaims{
			ExpiresAt: tokenExpireDuration.Unix(), //token的有效期
			IssuedAt:  time.Now().Unix(),          //token发放的时间
			Issuer:    "blog-back",                //作者
			Subject:   "sql token",                //主题
		},
	})
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(CustomSecret) //根据前面自定义的Jwt秘钥生成token
	if err != nil {
		//返回生成的错误
		return "", err
	}
	tokenString = "Bearer " + tokenString
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return tokenString, nil
}

func ParseToken(t string) (*jwt.Token, *CustomClaims, error) {
	claims := CustomClaims{}
	token, err := jwt.ParseWithClaims(t, &claims, func(token *jwt.Token) (interface{}, error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, nil, err
	}
	return token, &claims, nil
}
