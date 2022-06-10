package main

import "fmt"

type shape interface {
	accept(*visitor)
}

type Circle struct {
}

func (visitable *Circle) accept(visitor *visitor) {
	(*visitor).visitCircle(visitable)
}

type Rectangle struct {
}

func (visitable *Rectangle) accept(visitor *visitor) {
	(*visitor).visitRectangle(visitable)
}

type visitor interface {
	visitCircle(*Circle)
	visitRectangle(*Rectangle)
}

type areaVisitor struct {
}

func (visitor *areaVisitor) visitCircle(circle *Circle) {
	fmt.Println("visit circle area...")
}

func (visitor *areaVisitor) visitRectangle(rectangle *Rectangle) {
	fmt.Println("visit rectangle area...")
}

type perimeterVisitor struct {
}

func (visitor *perimeterVisitor) visitCircle(circle *Circle) {
	fmt.Println("visit circle perimeter...")
}

func (visitor *perimeterVisitor) visitRectangle(rectangle *Rectangle) {
	fmt.Println("visit rectangle perimeter...")
}
