package main

func main() {
	var arr0 = [2]iTarget{}
	arr0[0] = target{}
	arr0[1] = &adapter{
		adaptee{},
	}

	arr0[0].do()
	arr0[1].do()
}
