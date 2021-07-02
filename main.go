package main

// import "fmt"

// import "time"
import (
	"fmt"
	"time"
)

type point struct {
	x, y int
}

func main() {
	fmt.Println("Hello Go!")
	time.Sleep(1 * time.Second)
	p := point{1, 5}
	fmt.Printf(" %T\n %v\n %+v\n %#v\n", p, p, p, p)
}
