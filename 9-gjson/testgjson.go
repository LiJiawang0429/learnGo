package main

import (
	"fmt"

	"github.com/tidwall/gjson" //方法示例：https://github.com/tidwall/gjson
)

const json = `{
	"name": {"first": "Tom", "last": "Anderson"},
	"age":37,
	"children": ["Sara","Alex","Jack"],
	"fav.movie": "Deer Hunter",
	"friends": [
	  {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
	  {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
	  {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
	]
  }`

func main() {
	value := gjson.Get(json, "name.last")
	child := gjson.Get(json, "children")
	friend := gjson.Get(json, "friends.#.first")
	age := gjson.Get(json, "friends.#(age>45).age")

	println(value.String())
	println(child.String())
	println(friend.String())
	fmt.Println(age.Int())
	fmt.Printf("typeofbjson=%T\n",json)
}
