package main

import "fmt"

func main() {
	var factory OperatorFactory
	factory = &PlusOperatorFactory{}
	operator := factory.Create()
	operator.SetA(2)
	operator.SetB(1)
	res := operator.Operate()
	if res == 3 {
		fmt.Printf("plus success\n")
	} else {
		fmt.Printf("plus failed\n")
	}

	factory = &MinusOperatorFactory{}
	operator = factory.Create()
	operator.SetA(2)
	operator.SetB(1)
	res = operator.Operate()
	if res == 1 {
		fmt.Printf("minus success\n")
	} else {
		fmt.Printf("minus failed\n")
	}

}
