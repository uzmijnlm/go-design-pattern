package main

import "fmt"

func main() {
	var sub Subject
	sub = &Proxy{}

	res := sub.Do()

	fmt.Println(res)
}
