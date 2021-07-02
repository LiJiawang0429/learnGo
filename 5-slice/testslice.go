package main

import (
	"fmt"
)

func main() {
	myarray := []int{1, 2, 3, 4}
	myarray[3] = 500
	for _, value := range myarray {
		fmt.Println(value)
	}

	fmt.Println("---------------")
	slice := []int{1, 2, 3}

	var slice1 []int

	if slice1 == nil {
		fmt.Println("slice1为空！")
	} else {
		fmt.Println("slice1有空间!")
	}

	slice1 = make([]int, 3)

	fmt.Printf("len = %d,slice=%v\n", len(slice1), slice)

	if slice1 == nil {
		fmt.Println("slice1为空！")
	} else {
		fmt.Println("slice1有空间!")
	}
	fmt.Println("--------------------------------------")

	var slice2 = make([]int, 5, 7)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(slice2), cap(slice2), slice2)

	slice2 = append(slice2, 2)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(slice2), cap(slice2), slice2)

	slice2 = append(slice2, 8)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(slice2), cap(slice2), slice2)

	slice2 = append(slice2, 7)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(slice2), cap(slice2), slice2)

	fmt.Println("--------------------------------------")
	s1 := slice[0:2]
	fmt.Println(s1)
	copy(s1, slice2)
	fmt.Println(s1)

}
