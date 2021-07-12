package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i =%d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	//创建一个goroutine 执行newTask()
	go newTask()
	fmt.Printf("exit\n")

	// i := 0
	// for {
	// 	i++
	// 	fmt.Printf("main Goroutine : i =%d\n", i)
	// 	time.Sleep(1 * time.Second)
	// }

}
