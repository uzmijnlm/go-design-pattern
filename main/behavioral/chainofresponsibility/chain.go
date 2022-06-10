package main

import "fmt"

type Slot interface {
	check()
	setNext(*Slot)
}

type AbstractSlot struct {
	next *Slot
}

func (slot *AbstractSlot) setNext(next *Slot) {
	slot.next = next
}

func (slot *AbstractSlot) check() {
	fmt.Println("unimplemented")
}

type ParamSlot struct {
	AbstractSlot
}

func (slot *ParamSlot) check() {
	fmt.Println("ParamSlot check...")
	if slot.next != nil {
		(*slot.next).check()
	}
}

type DegradeSlot struct {
	AbstractSlot
}

func (slot *DegradeSlot) check() {
	fmt.Println("DegradeSlot check...")
	if slot.next != nil {
		(*slot.next).check()
	}
}

type SystemSlot struct {
	AbstractSlot
}

func (slot *SystemSlot) check() {
	fmt.Println("SystemSlot check...")
	if slot.next != nil {
		(*slot.next).check()
	}
}
