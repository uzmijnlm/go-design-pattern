package main

func main() {
	var shape Shape
	shape = &Rectangle{
		color: &Red{},
	}
	shape.print()

	shape = &Circle{
		color: &Green{},
	}
	shape.print()
}
