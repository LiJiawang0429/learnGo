package main

import "fmt"

type tree struct {
	Value          int
	LChild, RChild *tree
}

func Sort(values []int) {
	var root *tree
	// root := &tree{}//创建结构体变量
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	fmt.Println(values)
}

func add(r *tree, v int) *tree {
	if r == nil {
		r = new(tree)
		r.Value = v
		return r
	}
	if v < r.Value {
		r.LChild = add(r.LChild, v)

	}
	if v > r.Value {
		r.RChild = add(r.RChild, v)
	}
	fmt.Println(r)
	return r
}

func appendValues(values []int, r *tree) []int {
	if r != nil {
		values = appendValues(values, r.LChild)
		fmt.Println(values)
		values = append(values, r.Value)
		values = appendValues(values, r.RChild)
	}
	return values
}

func main() {
	value := []int{25, 54, 5, 59, 61, 12, 36, 98, 78}
	Sort(value)
}
