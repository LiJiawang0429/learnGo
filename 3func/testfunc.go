package main

import (
	"fmt"
)

func test(aa string, bb int) (r1 int, r2 string) {
	fmt.Println("aa=", aa)
	fmt.Println("bb=", bb)
	r1 = 200
	r2 = "dasdf"

	return
}
func test1(aa string, bb int) int {
	fmt.Println("aa=", aa)
	fmt.Println("bb=", bb)

	return 999
}

func main() {
	a, b := test("asd", 888)
	fmt.Println("a=", a, "\nb=", b)

	a = test1("asd", 888) //重新赋值不需要“：”
	fmt.Println("a=", a)
}
