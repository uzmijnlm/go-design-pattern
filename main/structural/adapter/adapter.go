package main

type iTarget interface {
	do()
}

type target struct {
}

func (target) do() {

}

type adaptee struct {
}

func (*adaptee) specialDo() {

}

type adapter struct {
	adaptee
}

func (a *adapter) do() {
	a.adaptee.specialDo()
}
