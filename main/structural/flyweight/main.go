package main

import "fmt"

func main() {
	viewer1 := NewImageViewer("image1.png")
	viewer2 := NewImageViewer("image1.png")

	if viewer1.image != viewer2.image {
		fmt.Println("error...")
	}
}
