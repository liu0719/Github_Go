package exec

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (m *Monster) Store() bool {
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("marshal err:", err)
		return false
	}
	err = ioutil.WriteFile("E:/go/src/go_study/unit_testing/exec/text.json", data, 0666)
	if err != nil {
		fmt.Println("writer err:", err)
		return false
	}
	return true
}
func (m *Monster) ReStore() bool {
	data, err := ioutil.ReadFile("E:/go/src/go_study/unit_testing/exec/text.json")
	if err != nil {
		fmt.Println("读取文件错误:", err)
		return false
	}
	err = json.Unmarshal(data, m)
	if err != nil {
		fmt.Println("反序列化错误:", err)
		return false
	}
	if m.Name != "孙悟空" {
		fmt.Println("名字错误。。")
		return false
	}
	return true
}
