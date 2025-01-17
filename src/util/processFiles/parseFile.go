package processFiles

import (
	"fmt"
	"html/template"
)

type ParseFile struct {
	urls []string
}

func NewParseFile(urls []string) *ParseFile {
	return &ParseFile{
		urls: urls,
	}
}

func recoverPanicOfTemplate() {
	if r := recover(); r != nil {
		fmt.Println("模板解析错误")
		return
	}
}

// ParseSingleFile 解析指定单独模板文件
func (f *ParseFile) ParseSingleFile(url string) *template.Template {
	defer recoverPanicOfTemplate()
	return template.Must(template.ParseFiles(url))
}

// ParseMultipleFile 解析指定多个模板文件
func (f *ParseFile) ParseMultipleFile() *template.Template {
	defer recoverPanicOfTemplate()
	return template.Must(template.ParseFiles(f.urls...))
}

// ParseFilesDefault 解析后台默认模板文件
func (f *ParseFile) ParseFilesDefault() *template.Template {
	defer recoverPanicOfTemplate()
	f.urls = []string{
		"views/pages/public/public_header.html",
		"views/pages/public/public_sidebar.html",
		"views/pages/users/list.html",
		"views/pages/users/info.html",
		"views/index.html",
	}

	return f.ParseMultipleFile()
}

// ParseFilesLogin  解析登录模板文件
func (f *ParseFile) ParseFilesLogin() *template.Template {
	return f.ParseSingleFile("views/pages/users/login.html")
}

// ParseFilesRegister  解析注册模板文件
func (f *ParseFile) ParseFilesRegister() *template.Template {
	return f.ParseSingleFile("views/pages/users/register.html")
}
