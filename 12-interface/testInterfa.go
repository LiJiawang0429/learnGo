package main

import "fmt"

type book struct {
	bookName string
}

//interface{}万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	//interface{}要如何区分 此时引用的底层数据类型
	//给 interface{}提供“类型断言”机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string")
	} else {
		fmt.Println("arg is string type, value =", value)
		fmt.Printf("value type is %T\n", value)

	}

}

func main() {
	book := book{"幸福来敲门"}
	myFunc(book)

	bookname := "prison break"
	myFunc(bookname)
}
