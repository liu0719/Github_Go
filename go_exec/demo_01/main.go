package main

import (
	"fmt"
	"go_exec/demo_01/utils"
	"os"
	"sync"
	"time"
)

var (
	//锁
	m sync.Mutex
	//主线程等待
	w sync.WaitGroup
)

func write(j int) error {
	// 枷锁
	// m.Lock()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()

	file, err := os.OpenFile("demo.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return err
	}
	uuid := utils.CreateUUID()
	// w.Add(1)
	for i := 0; i < 99; i++ {
		// w.Add(1)
		_, err := file.WriteString(uuid + "---" + fmt.Sprintf("%v====%v\n", i, j))
		if err != nil {
			fmt.Println(err)
		}
	}
	// w.Done()

	// m.Unlock()
	// w.Done()
	return err
}
func main() {
	start := time.Now().UnixNano()
	for i := 0; i < 200; i++ {
		// w.Add(1)
		write(i)
	}
	// w.Wait()
	end := time.Now().UnixNano()
	fmt.Println("共用时：", end-start)
}
