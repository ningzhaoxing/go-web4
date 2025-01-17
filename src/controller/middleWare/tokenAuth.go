package middleWare

import (
	"context"
	"net/http"
	token2 "userManageSystem-blog/src/util/token"
)

func AuthMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")

		if err != nil || cookie == nil {
			http.Redirect(w, r, "/user/login", http.StatusFound)
		}

		tokenString := cookie.Value

		//验证通过，提取有效部分（除去Bearer)
		tokenString = tokenString[7:] //截取字符
		//解析token
		token, _, err := token2.ParseToken(tokenString)

		if err != nil {
			http.Redirect(w, r, "/user/login", http.StatusFound)
		}
		// 将token.Claims存储到请求上下文
		ctx := context.WithValue(r.Context(), "claim", token.Claims)
		r = r.WithContext(ctx)
		handler(w, r)
	}
}
