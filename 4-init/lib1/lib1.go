package lib1

import (
	"fmt"
)

// 当前包的api
func Lib1test() {
	fmt.Println("lib1test()...")
}
func init() {
	fmt.Println("lib1.init()...")
}
