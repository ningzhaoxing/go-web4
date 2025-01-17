package customBind

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
)

type CustomBind struct {
	r *http.Request
}

func NewCustomBind(r *http.Request) *CustomBind {
	return &CustomBind{
		r: r,
	}
}

// BindQuery 将query数据绑定到结构体
func (lb *CustomBind) BindQuery(obj any) error {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return errors.New("请传递指针")
	}

	err := binding("query", obj, lb)
	if err != nil {
		return err
	}
	return nil
}

// BindForm 将form表单数据绑定到结构体
func (lb *CustomBind) BindForm(obj any) error {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return errors.New("请传递指针")
	}

	err := binding("form", obj, lb)
	if err != nil {
		return err
	}
	return nil
}

func binding(queryTyp string, obj any, lb *CustomBind) error {
	val := reflect.ValueOf(obj).Elem()
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("json")

		// 跳过未导出字段
		if tag == "" {
			continue
		}

		var value string
		if queryTyp == "form" {
			value = lb.r.FormValue(tag)
		} else if queryTyp == "query" {
			value = lb.r.URL.Query().Get(tag)
		}

		// 跳过未传入的参数字段
		if value == "" {
			continue
		}

		fieldVal := val.Field(i)
		if fieldVal.Kind() == reflect.String {
			fieldVal.SetString(value)
		} else if fieldVal.Kind() == reflect.Int {
			value, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("数据绑定int类型转化错误")
				return err
			}
			fieldVal.SetInt(int64(value))
		} else if fieldVal.Kind() == reflect.Bool {
			value, err := strconv.ParseBool(value)
			if err != nil {
				fmt.Println("数据绑定bool类型转化错误")
				return err
			}
			fieldVal.SetBool(value)
		}
		// 11111
	}
	return nil
}
