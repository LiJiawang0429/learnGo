package main

import (
	"fmt"
)

func printarray(myarray [4]int) {
	for index, value := range myarray {
		fmt.Println(index, value)
	}
}

func main() {
	var myarray1 [10]int
	myarray2 := [10]int{1, 2, 3, 4, 5}
	myarray3 := [4]int{1, 2, 3, 4}
	for i := 0; i < len(myarray1); i++ {
		fmt.Println(myarray1[i])
	}

	for index, value := range myarray2 {
		fmt.Println("index=", index, "value=", value)
	}

	printarray(myarray3)
}
