package main

import "fmt"

type Builder interface {
	step1()
	step2()
}

type Builder1 struct {
}

func (Builder1) step1() {
	fmt.Println("Builder1 step1")
}

func (Builder1) step2() {
	fmt.Println("Builder1 step2")
}

type Builder2 struct {
}

func (Builder2) step1() {
	fmt.Println("Builder2 step1")
}

func (Builder2) step2() {
	fmt.Println("Builder2 step2")
}

type Director struct {
	builder Builder
}

func (d Director) construct() {
	d.builder.step1()
	d.builder.step2()
}
