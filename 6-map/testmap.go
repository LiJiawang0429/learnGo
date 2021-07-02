package main

import (
	"fmt"
)

func printmap(citymap map[string]string) {

	for key, value := range citymap {
		fmt.Println("key=", key, " value=", value)
	}
}

func Changevalue(citymap map[string]string) {
	citymap["english"] = "London"

}
func main() {
	citymap := make(map[string]string)

	citymap["China"] = "Beijing"
	citymap["japan"] = "tokyo"
	citymap["usa"] = "newyork"

	printmap(citymap)

	delete(citymap, "japan")

	citymap["usa"] = "DC"
	Changevalue(citymap)
	fmt.Println("--------------------------------------")

	printmap(citymap)

}
