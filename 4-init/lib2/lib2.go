package lib2

import (
	"fmt"
)

// 当前包的api
func Lib2test() {
	fmt.Println("lib2test()...")
}

func init() {
	fmt.Println("lib2.init()....")
}
