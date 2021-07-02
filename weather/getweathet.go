package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
	// "github.com/tidwall/gjson" //方法示例：https://github.com/tidwall/gjson
)

func main() {
	//https://www.tianqiapi.com/free/day?appid=88366117&appsecret=yfW1E3HM
	url := "https://www.tianqiapi.com/free/day?appid=88366117&appsecret=yfW1E3HM"
	get_url(url)
	time.Sleep(1 * time.Millisecond)

}
func get_url(url string) {
	//fmt.Println("----------", url, "----------------")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http get error.")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
	}
	//fmt.Printf("typeofbody=%T\n",body)
	json := string(body)
	//fmt.Println(src)
	cityName := gjson.Get(json, "city")
	UpdateTime := gjson.Get(json, "update_time")
	cityWeather := gjson.Get(json, "wea")
	temperature := gjson.Get(json, "tem")
	win := gjson.Get(json, "win")
	wins := gjson.Get(json, "win_speed")
	air := gjson.Get(json, "air")
	fmt.Println("---------------------------------")
	fmt.Println("----------本地天气情况------------")
	fmt.Printf("城市：%s         时间：%s\n", cityName, UpdateTime)
	fmt.Printf("当前天气：%s     此刻气温：%s℃\n", cityWeather, temperature)
	fmt.Printf("此刻风向：%3s    风力：%s\n", win, wins)
	fmt.Printf("空气指数：%s\n", air)
	fmt.Println("---------------------------------")

}
