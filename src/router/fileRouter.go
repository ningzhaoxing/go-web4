package router

import (
	"github.com/gorilla/mux"
	"userManageSystem-blog/src/controller/file"
	"userManageSystem-blog/src/controller/middleWare"
)

func fileRouter(r *mux.Router) {
	r.HandleFunc("/upload/userHeadPhoto", middleWare.AuthMiddleware(file.UploadHeadPhotoController))
}
