package main

import "fmt"

func main() {
	subject := &Subject{context: "0"}
	var observer1 Observer
	observer1 = &Reader{
		name:    "reader1",
		context: "0",
	}
	subject.addObserver(&observer1)
	var observer2 Observer
	observer2 = &Reader{
		name:    "reader2",
		context: "0",
	}
	subject.addObserver(&observer2)

	subject.updateSubject("1")

	fmt.Println()
}
