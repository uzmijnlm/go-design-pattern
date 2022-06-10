package main

import (
	"fmt"
)

func main() {
	api := NewAPI(1)
	res := api.Say("Tom")
	fmt.Printf(res)
}
