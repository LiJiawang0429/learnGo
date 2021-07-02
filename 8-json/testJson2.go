package main

import (
	"encoding/json"
	"fmt"
)

type Fruit struct { //制作结构
	Describe string `json:"describe"`
	Icon     string `json:"icon"`
	Name     string `json:"name"`
}

type FruitGroup struct { //嵌套制作结构
	FirstFruit      *Fruit   `json:"first fruit"`  //指针，指向引用对象；如果不用指针，只是值复制
	SecondFruit     *Fruit   `json:"second fruit"` //指针，指向引用对象；如果不用指针，只是值复制
	THreeFruitArray []string `json:"three fruit array"`
}

func main() {
	//转为Json类型
	var jsonBlob = []byte(`{
    "first fruit": {
        "describe": "an apple",
        "icon": "appleIcon",
        "name": "apple"
    },
    "second fruit": {
        "describe": "an orange",
        "icon": "appleIcon",
        "name": "orange"
    },
    "three fruit array": [
        "eat 0",
        "eat 1",
        "eat 2",
        "eat 3"
    ]}`)

	var fruitGroup FruitGroup //定义结构体

	err := json.Unmarshal(jsonBlob, &fruitGroup) //解析jsonBlob，并使指针firstGroup指向它
	if err != nil {
		fmt.Println("error:", err) //检测
	}
	fmt.Printf("%+v\n", fruitGroup) //输出go结构变量
	fmt.Printf("%+v\n", fruitGroup.FirstFruit)
	fmt.Printf("%+v\n", fruitGroup.SecondFruit)
}
