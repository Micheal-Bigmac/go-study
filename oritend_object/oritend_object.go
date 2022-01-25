package oritend_object

import "fmt"

type Animal struct {
	Colour string
}

type Dog struct {
	Animal
	Id   int
	Name string
	Age  int
}

type Cat struct {
	Animal
	Id   int
	Name string
	Age  int
}

func (a *Dog) Eat() string {
	fmt.Println("Dog chi chichi ")
	return "dog chi chichi"
}

func (a *Dog) Run() string {
	fmt.Println("ID : ", a.Id, " Dog is running ")
	return "Dog is running "
}

func (a *Cat) Eat() string {
	fmt.Println("cat chi chichi ")
	return " cat  chi chichi"
}

func (a *Cat) Run() string {
	fmt.Println("ID : ", a.Id, "Cat is running ")
	return "cat is running "
}
