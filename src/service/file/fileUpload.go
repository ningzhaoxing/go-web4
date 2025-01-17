package file

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type UploadFileInterface interface {
	GetUploadFile() multipart.File // 文件
	GetFilePath() string           // 文件路径
}

type UploadFile struct {
	ut UploadFileInterface
}

func NewUploadFile(ut UploadFileInterface) *UploadFile {
	return &UploadFile{
		ut: ut,
	}
}

// Upload 实现文件上传
func (uf *UploadFile) Upload() (string, error) {
	file := uf.ut.GetUploadFile()
	filePath := uf.ut.GetFilePath()
	filePath = fmt.Sprintf("uploads/%d/%s/%d/%s", time.Now().Year(), time.Now().Month(), time.Now().Day(), filePath)

	dir := filepath.Dir(filePath)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return "", err
	}

	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	filePath = fmt.Sprintf("%s%s", "http://localhost:8899/", filePath)
	return filePath, nil
}
