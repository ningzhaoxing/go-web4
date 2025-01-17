package response

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"userManageSystem-blog/src/pkg/globals"
)

type Response struct {
	w http.ResponseWriter
}

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{w: w}
}

func (r *Response) Success(data interface{}, t *template.Template, temp string) {
	if data == nil {
		data = struct{}{}
	}

	err := t.ExecuteTemplate(r.w, "index", Appdata{
		Code: globals.CodeSuccess,
		Msg:  "",
		Data: data,
		Temp: temp,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func (r *Response) Fail(err error, t *template.Template, name string) {
	e := err.Error()
	_ = t.ExecuteTemplate(r.w, name, &Appdata{
		Code: globals.CodeFailed,
		Msg:  e,
		Data: "",
	})
}

func (r *Response) HttpSuccess(msg string, data interface{}) {
	if data == nil {
		data = struct{}{}
	}
	r.JSON(http.StatusOK, &Appdata{
		Code: globals.CodeSuccess,
		Msg:  msg,
		Data: data,
	})
}

func (r *Response) HttpFail(err error) {
	e := err.Error()
	r.JSON(http.StatusOK, &Appdata{
		Code: globals.CodeFailed,
		Msg:  e,
		Data: "",
	})
}

func (r *Response) JSON(code int, data *Appdata) {
	r.w.Header().Set("Content-Type", "application/json")
	r.w.WriteHeader(code)
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(r.w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = r.w.Write(jsonData)
	if err != nil {
		http.Error(r.w, err.Error(), http.StatusInternalServerError)
		return
	}
}
