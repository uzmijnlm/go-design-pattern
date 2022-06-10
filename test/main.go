package main

import "fmt"

type Animal interface {
	eat()
}

type Dog struct {
}

type Cat struct {
}

func (*Cat) eat() {
	fmt.Println("cat eat")
}

func (*Dog) eat() {
	fmt.Println("dog eat")
}

type Pet struct {
	*Cat
	*Dog
}

func (pet *Pet) eat() {
	fmt.Printf("%p\n", pet)
	fmt.Println("pet eat")
}

func main() {
	fmt.Println("test")
	pet := new(Pet)
	fmt.Printf("%p\n", pet)
	pet.eat()
}
