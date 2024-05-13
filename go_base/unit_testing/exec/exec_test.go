package exec

import (
	"testing"
)

func TestStore(t *testing.T) {
	m1 := &Monster{
		Name:  "孙悟空",
		Age:   1000,
		Skill: "七十二变",
	}
	res := m1.Store()
	if !res {
		t.Fatalf("序列化错误")
	}
	t.Logf("序列化成功")
}

func TestReStore(t *testing.T) {
	mon := Monster{}
	flag := mon.ReStore()
	if !flag {
		t.Fatalf("反序列化错误")
	}
	t.Logf("反序列化成功")
}
