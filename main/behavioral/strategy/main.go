package main

import "fmt"

func main() {
	payment := NewPayment("Ada", "", 123, &Cash{})
	payment.Pay()

	fmt.Println()

	payment = NewPayment("Bob", "", 123, &Bank{})
	payment.Pay()
}
