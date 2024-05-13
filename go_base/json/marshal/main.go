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

//序列化结构体时，字段首字母一定要大写,挎包在别的文件中无法使用。
func testStruct() {
	//演示
	// 结构体序列化
	monster := Monster{
		Name:     "牛魔王",
		Age:      300,
		Brithday: "2011 - 1 - 11",
		Sal:      900.0,
		Skill:    "牛魔拳",
	}
	//将monster序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化失败：", err)
	}
	fmt.Println(string(data))
	err1 := ioutil.WriteFile("./json/serialization/textStruct.json", data, 0666)
	if err != nil {
		fmt.Println(err1)
	}
}

func testMap() {
	//定义一个map
	a := make(map[string]map[string]string)
	a["001"] = make(map[string]string)
	a["001"]["name"] = "红孩儿"
	a["001"]["age"] = "100"
	a["001"]["sal"] = "800"
	a["001"]["adderss"] = "火云洞"
	a["002"] = make(map[string]string)
	a["002"]["skill"] = "三味真火"
	a["002"]["name"] = "铁扇公主"
	a["002"]["age"] = "300"
	a["002"]["sal"] = "8000"
	a["002"]["adderss"] = "未知"
	a["002"]["skill"] = "芭蕉扇"
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	_ = ioutil.WriteFile("./json/serialization/textMap.json", data, 0666)
	fmt.Println(string(data))
}

//切片序列化，[]map[string]interface{}
func testSlice() {
	slice := make([]map[string]interface{}, 0)
	s1 := make(map[string]interface{})
	s1["name"] = "jack"
	s1["age"] = 20
	s1["adderss"] = &Monster{
		Name:     "孙悟空",
		Age:      1000,
		Brithday: "nil",
		Sal:      9999,
		Skill:    "七十二变",
	}
	slice = append(slice, s1)
	s2 := make(map[string]interface{})
	s2["name"] = "mary"
	s2["age"] = 25
	s2["adderss"] = &Monster{
		Name:     "白骨经",
		Age:      100,
		Brithday: "111",
		Sal:      99,
		Skill:    "九阴白骨爪",
	}
	slice = append(slice, s2)
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("err")
	}
	_ = ioutil.WriteFile("./json/serialization/textSlice.json", data, 0666)
	fmt.Println(string(data))

}

//对基本类型序列化不大
func testFloat64() {
	num1 := 23456.7980
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("_____________________________________")
	fmt.Println(string(data))
}
func main() {
	testStruct()  //结构体序列化
	testMap()     //map序列化
	testSlice()   //切片序列化
	testFloat64() //基本数据类型序列化
}
