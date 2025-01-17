package middleWare

import (
	"net/http"
	"userManageSystem-blog/src/pkg/response"
)

func CORSMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 允许所有源
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// 允许所有方法
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		// 允许的头信息
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		// 允许浏览器获取响应的头部信息
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		// 允许携带凭证（如cookies）
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		res := response.NewResponse(w)
		if r.Method == "OPTIONS" {
			res.HttpSuccess("", struct{}{})
			return
		}
		handler(w, r)
	}
}
