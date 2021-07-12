package main

import "fmt"

type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
}

func (d *Cat) Sleep() {
	fmt.Println("Cat is sleep")
}

func (c Cat) GetColor() string {
	return c.color
}

func (c Cat) GetType() string {
	return "cat"
}

type Dog struct {
	color string
}

func (d *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}

func (d Dog) GetColor() string {
	return d.color
}

func (d Dog) GetType() string {
	return "dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep() //多态
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("kind = ", animal.GetType())
}

func main() {
	cat := Cat{"white"}
	dog := Dog{"Yellow"}

	showAnimal(&cat)
	showAnimal(&dog)

	// var animal AnimalIF
	// animal = &Dog{"Yellow"}
	// animal.Sleep()
	// animal.GetType()
	// animal.GetColor()

	// animal = &Cat{"white"}
	// animal.Sleep()
	// animal.GetType()
	// animal.GetColor()
}
