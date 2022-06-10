package main

func main() {
	circle := &Circle{}
	rectangle := &Rectangle{}

	var visitor1, visitor2 visitor
	visitor1 = &areaVisitor{}
	visitor2 = &perimeterVisitor{}

	circle.accept(&visitor1)
	circle.accept(&visitor2)
	rectangle.accept(&visitor1)
	rectangle.accept(&visitor2)

}
