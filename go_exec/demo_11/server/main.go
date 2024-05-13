package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "192.168.1.4:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go process(c)
	}

}

func process(conn net.Conn) {
	// 函数执行完之后关闭连接
	defer conn.Close()
	// 输出主函数传递的conn可以发现属于*TCPConn类型, *TCPConn类型那么就可以调用*TCPConn相关类型的方法, 其中可以调用read()方法读取tcp连接中的数据
	fmt.Printf("有人连接: %T\n", conn)
	for {
		var buf [128]byte
		// 将tcp连接读取到的数据读取到byte数组中, 返回读取到的byte的数目
		n, err := conn.Read(buf[:])
		fmt.Println(n)
		if err != nil {
			// 从客户端读取数据的过程中发生错误
			fmt.Println("read from client failed, err:", err)
			break
		}
		// recvStr := string(buf[:n])
		fmt.Println("服务端收到客户端发来的数据：", buf[:n])
	}
}
