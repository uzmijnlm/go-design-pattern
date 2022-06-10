package main

func main() {
	var builder Builder
	var director *Director
	builder = &Builder1{}
	director = &Director{builder: builder}
	director.construct()

	builder = &Builder2{}
	director = &Director{builder: builder}
	director.construct()
}
