package main

// import "fmt"

// import "time"
import (
	"fmt"
	"strconv"
	//"time"
	// "time"
)

// type point struct {
// 	x, y int
// }

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		// if i != j {
		// 	t := a[i]
		// 	a[i] = a[j]
		// 	a[j] = t
		// }
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a)
}

func twoSum(nums []int, target int) {
	var r [4]string
	k := 0
	for i := 0; i < len(nums); i++ {
		a := nums[i]
		for j := i + 1; j < len(nums); j++ {
			b := nums[j]
			sum := a + b
			if sum == target {
				r[k] = strconv.Itoa(i) + "," + strconv.Itoa(j)
				k += 1
				fmt.Println("ok", r)
			}
		}
	}
	fmt.Println(r)
}
func main() {
	arr := []int{2, 7, -6, 15}
	reverse(arr)
	t := 1
	twoSum(arr, t)

	// fmt.Println("result=", r)
	// fmt.Println("Hello Go!")
	// time.Sleep(1 * time.Second)
	// p := point{1, 5}
	// fmt.Printf(" %T\n %v\n %+v\n %#v\n", p, p, p, p)
}
