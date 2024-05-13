package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//os.Args获取命令行参数
	fmt.Println("命令行的参数有", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}

	//flag包解析命令行参数

	//定义几个变量，接收命令行参数值
	var user string
	var pwd string
	var host string
	var port int
	// &user 就是接收用户命令行中的 -u 后面的参数值
	// "u"就是-u指定参数
	// ""默认值
	// "用户名，默认为空"说明
	flag.StringVar(&user, "u", "", "用户名默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号")

	//这里有一个非常重要的操作，必须调用该方法
	flag.Parse()

	//输出结果
	fmt.Printf("user=%v,pwd=%v,host=%v,port=%v\n", user, pwd, host, port)
	//优点：可以不按照顺序进行输入，-u后面的是user的参数，
}
