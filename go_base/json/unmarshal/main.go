package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name     string  `json:"name"` //控制json序列化后的json的key值，反射机制
	Age      int     `json:"age"`
	Brithday string  `json:"brithday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

//演示将json字符串，反序列化成struct
func unmarshalstruct() {
	data, err := ioutil.ReadFile("./json/marshal/textStruct.json")
	if err != nil {
		fmt.Println(err)
	}
	//在开发中，是通过网络传输获取的

	m := Monster{}
	//第一个参数是要转的数据，要是byte数组的格式，第二个参数是接收的类型，不是引用类型传地址效率高
	_ = json.Unmarshal(data, &m)
	fmt.Println(m)
}

func unmarshalMap() {
	data, err := ioutil.ReadFile("./json/marshal/textMap.json")
	if err != nil {
		fmt.Println(err)
	}
	//反序列化
	//注意：反序列化map，不需要make分配内存，因为make操作被封装到unmarshal函数里了
	var m1 map[string]interface{}
	_ = json.Unmarshal(data, &m1)
	fmt.Println(m1)
}

func unmarshalSlice() {
	data, err := ioutil.ReadFile("./json/marshal/textSlice.json")
	if err != nil {
		fmt.Println(err)
	}
	var a []map[string]interface{}
	_ = json.Unmarshal(data, &a)
	fmt.Println(a)
}
func main() {
	unmarshalstruct()
	unmarshalMap()
	unmarshalSlice()
}
