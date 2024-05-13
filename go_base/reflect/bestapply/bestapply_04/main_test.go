package main

import (
	"reflect"
	"testing"
)

type user struct {
	UserId string
	Name   string
}

func TestReflectStruct(t *testing.T) {
	var (
		model *user
		st    reflect.Type
		elem  reflect.Value
	)
	st = reflect.TypeOf(model)                  //获取类型*user
	t.Log("reflect.TypeOf", st.Kind().String()) //ptr
	st = st.Elem()                              //st指向的类型
	elem = reflect.New(st)
	t.Log("reflect.New", elem.Kind().String()) //ptr
	t.Log("reflect.New.Elem", elem.Elem().Kind().String())
	//model就是创建的user结构体变量
	model = elem.Interface().(*user) //model是*user它的指向和elem是一样的
	elem = elem.Elem()               //取得elem指向的值
	elem.FieldByName("UserId").SetString("12345678")
	elem.FieldByName("Name").SetString("nickname")
	t.Log("model model.Name", model, model.Name)

}
