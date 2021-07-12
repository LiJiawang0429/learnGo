package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"Name" doc:"我的名字"`
	Sex  string
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Call() {
	fmt.Println("User is called..")
	fmt.Printf("%v\n", u)
}

func DoMethod(input interface{}) {
	//获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is :", inputType.Name())

	//获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is :", inputValue)

	//通过Type获取里面的字段
	//1.获取interface的reflect.Type，通过Type得到NumField，进行遍历
	//2.得到每一个field，数据类型
	//3.通过field有一个Interface（）方法得到 对应的Value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)

	}

	//通过type 获取里面的方法，调用
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

func DoTagMethod(Resume interface{}) {
	r := reflect.TypeOf(Resume).Elem()

	for i := 0; i < r.NumField(); i++ {
		taginfo := r.Field(i).Tag.Get("info")
		tagdoc := r.Field(i).Tag.Get("doc")
		fmt.Printf("info:%s\n", taginfo)
		fmt.Printf("doc:%s\n", tagdoc)
	}
}

func main() {
	// user := User{11, "极佳", 22}
	// DoMethod(user)
	var re resume
	DoTagMethod(&re)
}
