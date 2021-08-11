package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name      string
	Pid       int
	Intime    time.Time
	ManagerID int
}


func main() {

	var p Employee
	name := &p.Name
	*name = "Wang" + *name
	p.Pid = 15264
	p.ManagerID = 10546
	fmt.Println(p)
}
