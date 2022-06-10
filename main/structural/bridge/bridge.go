package main

import "fmt"

type Shape interface {
	print()
}

type Color interface {
	printDetail()
}

type Red struct {
}

func (*Red) printDetail() {
	fmt.Print("红色的")
}

type Green struct {
}

func (*Green) printDetail() {
	fmt.Print("绿色的")
}

type Circle struct {
	color Color
}

func (circle *Circle) print() {
	circle.color.printDetail()
	fmt.Println("圆形")
}

type Rectangle struct {
	color Color
}

func (rectangle *Rectangle) print() {
	rectangle.color.printDetail()
	fmt.Println("长方形")
}
