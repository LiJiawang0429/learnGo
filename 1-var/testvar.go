package main

import (
	"fmt"
	"os"
)

// var glo int = 666

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		fmt.Println("dasdadasd", os.Args[i])
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	// //可用作声明全局变量
	// fmt.Println("glo=", glo)

	// var a string
	// fmt.Println(a)
	// var b int = 100
	// fmt.Println(b)
	// var c int = 100
	// fmt.Println(c)
	// var b1 = "dasffa"
	// fmt.Printf("b1 = %s ,type of b1 = %T\n", b1, b1)
	// fmt.Printf("\n\n")

	// //函数内常用方法，不支持全局变量
	// d := 100
	// // fmt.Printf("d =",d)  //Printf用于格式化输出
	// fmt.Println("d = ", d) //Println直接输出值
	// fmt.Printf("type of d = %T\n", d)

	// e := "ssdf"
	// fmt.Println("e = ", e) //Println直接输出值
	// fmt.Printf("type of e= %T\n", e)

	// f := 54.897
	// fmt.Println("f = ", f) //Println直接输出值
	// fmt.Printf("type of f = %T\n", f)

	// //多变量同时声明
	// var (
	// 	aa string = "sdf"
	// 	bb bool   = false
	// )
	// var cc, dd = "erer", 999
	// fmt.Println("aa,bb,cc,dd:", aa, bb, cc, dd)
	// fmt.Println("aa,bb,cc,dd:"+aa, bb, cc, dd)
	// fmt.Println("ds")
	// fmt.Printf("快捷缩写打印%d\n", d)
}
