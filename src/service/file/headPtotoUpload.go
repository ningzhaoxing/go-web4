package file

import (
	"fmt"
	"mime/multipart"
	"userManageSystem-blog/src/dao/db/userDb"
	"userManageSystem-blog/src/model/params/sql"
	"userManageSystem-blog/src/model/user"
	"userManageSystem-blog/src/pkg/globals"
)

type HeadPhotoUpload struct {
	err      error
	u        user.User
	file     multipart.File
	filePath string
	appCtx   *globals.AppCtx
}

func (h *HeadPhotoUpload) GetUploadFile() multipart.File {
	return h.file
}

func (h *HeadPhotoUpload) GetFilePath() string {
	return fmt.Sprintf("%s/%s", h.u.Id, h.filePath)
}

func (h *HeadPhotoUpload) Upload() error {
	h.SaveFileInLocal().SaveFileUrlInDb()
	return h.err
}

func NewHeadPhotoUpload(u user.User, file multipart.File, fileName string, appCtx *globals.AppCtx) *HeadPhotoUpload {
	return &HeadPhotoUpload{
		err:      nil,
		u:        u,
		file:     file,
		filePath: fileName,
		appCtx:   appCtx,
	}
}

// SaveFileInLocal 将文件保存到本地
func (h *HeadPhotoUpload) SaveFileInLocal() *HeadPhotoUpload {
	h.filePath, h.err = NewUploadFile(h).Upload()
	return h
}

// SaveFileUrlInDb 将图片url保存到数据库
func (h *HeadPhotoUpload) SaveFileUrlInDb() *HeadPhotoUpload {
	if h.err != nil {
		return h
	}

	h.err = userDb.UpdateUserHeadPhoto(sql.UserSqlParam{
		Db: h.appCtx.GetDb(),
		User: user.User{
			Id:        h.u.Id,
			HeadPhoto: h.filePath,
		},
	})
	return nil
}
