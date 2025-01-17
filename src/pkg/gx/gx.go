package gx

import (
	"net/http"
	"userManageSystem-blog/src/pkg/response"
	"userManageSystem-blog/src/util/customBind"
	"userManageSystem-blog/src/util/processFiles"
)

type GContext struct {
	t    *processFiles.ParseFile
	bind *customBind.CustomBind
	res  *response.Response
}

type Options func(g *GContext)

func NewGx(op ...Options) *GContext {
	gx := new(GContext)
	for _, options := range op {
		options(gx)
	}
	return gx
}

func NewDefaultGx(w http.ResponseWriter, req *http.Request) *GContext {
	return NewGx(WithOptionBind(req), WithOptionResponse(w), WithOptionTemplate())
}

func WithOptionTemplate() Options {
	return func(g *GContext) {
		g.t = processFiles.NewParseFile(nil)
	}
}

func WithOptionBind(r *http.Request) Options {
	return func(g *GContext) {
		g.bind = customBind.NewCustomBind(r)
	}
}

func WithOptionResponse(w http.ResponseWriter) Options {
	return func(g *GContext) {
		g.res = response.NewResponse(w)
	}
}

func (g *GContext) GetTemplate() *processFiles.ParseFile {
	return g.t
}

func (g *GContext) GetBind() *customBind.CustomBind {
	return g.bind
}

func (g *GContext) GetResponse() *response.Response {
	return g.res
}
