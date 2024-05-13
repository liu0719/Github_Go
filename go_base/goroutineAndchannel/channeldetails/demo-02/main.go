package main

import (
	"fmt"
)

func main() {
	//使用select可以解决从管道取数据的阻塞问题
	intchan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intchan <- i
	}
	//定义一个管道，5个string
	stringchan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringchan <- "hello" + fmt.Sprintf("%d", i)
	}

	//传统的方法在遍历管道是会阻塞而导致阻塞，而导致死锁

	//问题在实际开发中，可能我们不确定什么时候关闭管道
	//使用select可以解决问题
lable:
	for {
		select {
		case v := <-intchan:
			//如果intchan一直没有关闭，不会一直阻塞而导致死锁
			//会自动到下一个case进行匹配
			fmt.Printf("从intchan 读取了数据%v\n", v)
		case v := <-stringchan:
			fmt.Printf("从stringchan中取出的数据%v\n", v)
		default:
			fmt.Println("都没啦")
			break lable
		}
	}
}
