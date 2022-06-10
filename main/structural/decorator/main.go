package main

import "fmt"

func main() {
	var coffee Drink = &Espresso{}
	fmt.Printf("%T", coffee)

	coffee = &Sugar{
		drink: coffee,
	}
	coffee = &Milk{
		drink: coffee,
	}
	res := coffee.getPrice()

	fmt.Printf("res %d\n", res)
	// Output:
	// res 80
}
