package cal

//引入go的testing框架包
import (
	"fmt"
	"testing"
)

/*
	命令行指令
	go test 如果运行正确，不输出日志，错误输出日志
	go test -v 不论正确与否，都输出日志
*/
//编写测试用例，测试addupper是否正确
func TestAddUpper(t *testing.T) { //传入的testing.T必须是指针，
	//调用
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("Addupper(10)执行错误,期望值=%v,实际值%v\n", 55, res)
	}
	t.Logf("Addupper(10) 执行正确...")
}

func TestHello(t *testing.T) {
	fmt.Println("hello被调用...")
}
