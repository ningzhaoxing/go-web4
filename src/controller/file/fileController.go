package file

import (
	"fmt"
	"net/http"
	"userManageSystem-blog/src/pkg/globals"
	"userManageSystem-blog/src/pkg/response"
	file2 "userManageSystem-blog/src/service/file"
	"userManageSystem-blog/src/util/token"
)

// UploadHeadPhotoController 上传头像
func UploadHeadPhotoController(w http.ResponseWriter, r *http.Request) {
	appCtx := globals.NewDefaultAppCtx()
	res := response.NewResponse(w)

	if err := r.ParseMultipartForm(10 << 20); err != nil { // 设置最大内存为10MB
		res.HttpFail(err)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Println("controller.fileController.FormFile() err=", err)
		res.HttpFail(err)
		return
	}
	defer file.Close()
	user, err := token.NewToken(r).GetUser(appCtx.GetDb())
	if err != nil {
		fmt.Println("controller.fileController.GetUser() err=", err)
		res.HttpFail(err)
		return
	}

	err = file2.NewHeadPhotoUpload(*user, file, header.Filename, appCtx).Upload()
	if err != nil {
		fmt.Println("controller.fileController.NewHeadPhotoUpload() err=", err)
		res.HttpFail(err)
		return
	}

	res.HttpSuccess("上传成功", struct{}{})
}
