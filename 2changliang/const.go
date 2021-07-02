package main

import (
	"fmt"
)
// 常量声明
const (
	// iota定义枚举类型，只能配合const使用
	a = iota   //每行累加
	b
	c
	d
)
const (
	aa, bb = iota + 1, iota + 2
	cc, dd
	ee, ff = iota * 2, iota * 3
	gg, hh
)

func main() {
	e := 100
	fmt.Println(e)
	fmt.Println("a = ", a, "b = ", b)
	fmt.Println("c = ", c, "d = ", d)

	fmt.Println("aa = ", aa, "bb = ", bb)
	fmt.Println("cc = ", cc, "dd = ", dd)
	fmt.Println("ee = ", ee, "ff = ", ff)
	fmt.Println("gg = ", gg, "hh = ", hh)
}
