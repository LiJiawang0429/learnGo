package main

import (
	"fmt"
)

type Hero struct { //首字母大写表示可以被外访问！
	Name  string
	Ad    int
	Level int
}

type SuperHero struct {
	Hero
	Attack int
}

// func (this Hero) Show() {  //(this Hero)表示当前方法绑定到此结构体
// 	fmt.Println("Name = ", this.Name)
// 	fmt.Println("Ad = ", this.Ad)
// 	fmt.Println("Level = ", this.Level)
// }

// func (this Hero) GetName() {
// 	fmt.Println("Name=", this.Name)
// }

// func (this Hero) SetName(newName string) {
// 	this.Name = newName
// }

func (this *Hero) Show() { //结构体调用加*，否则调用的仅为副本
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}

func (this *Hero) GetName() {
	fmt.Println("Name=", this.Name)
}

func (this *Hero) SetName(newName string) {
	this.Name = newName
}

func (this *SuperHero) Fly() {
	fmt.Println("SuperHero can fly()")
}

func main() {
	hero := Hero{Name: "Z3", Ad: 100, Level: 1}
	hero.GetName()
	hero.SetName("L4")
	hero.GetName()
	fmt.Println("——————————————————————————————————————")
	// s := SuperHero{Hero{"W5",78,22},99999}   //利用父类结构定义
	var s SuperHero //也可单独定义，逐个赋值
	s.Name = "Z6"
	s.Level = 100
	s.Ad = 0
	s.Attack = 1
	s.Show() //继承父类函数
	s.Fly()  //使用自身函数

}
