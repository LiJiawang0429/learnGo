package main

import (
	"fmt"
	"time"
	_ "time"
)

func main() {
	c := make(chan int, 3)
	// 用go创建一个形参为空返回值为空的一个函数
	go func() {
		defer fmt.Println("defer a")

		func() {
			defer fmt.Println("defer b")
			time.Sleep(2 * time.Second)
			for i := 0; i < 8; i++ {
				c <- i
				fmt.Println("子go程正在运行：len(c)=", len(c), ",cap(c)=", cap(c))
			}
			//关闭当前channel
			close(c)

			// 退出当前goroutine
			//runtime.Goexit()
			fmt.Println("b")
		}()

		fmt.Println("a")

	}()

	// 用go创建一个形参存在返回值存在的一个函数
	// go func(a, b int) bool {
	// 	fmt.Println(a, b)
	// 	return true
	// }(50, 100)

	for data := range c {
		fmt.Println("获取goroutine中的c：", data)
	}
	// fmt.Println("Main finished")
	// 死循环
	// for {
	// 	time.Sleep(1 * time.Second)
	// }
}
