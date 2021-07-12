package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Name   string   `json:"name"`
	Actors []string `json:"actors"`
	Year   int
	Price  int
}

func main() {

	movie := Movie{"西红柿首富", []string{"沈腾", "ljw"}, 2020, 50}

	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal err!!", err)
		return
	}
	fmt.Printf("jsonStr=%s\n", jsonStr)

	movieT := Movie{}

	err = json.Unmarshal(jsonStr, &movieT)
	if err != nil {
		fmt.Println("json Unmarshal error!", err)
		return
	}
	fmt.Printf("movie = %v\n", movieT)
}
