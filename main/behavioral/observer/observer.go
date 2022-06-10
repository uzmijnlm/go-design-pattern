package main

import "fmt"

type Observer interface {
	Update(*Subject)
}

type Reader struct {
	name    string
	context string
}

func (reader *Reader) Update(subject *Subject) {
	before := reader.context
	context := subject.context
	reader.context = context
	after := reader.context
	fmt.Println("reader name is " + reader.name + ". before: " + before + ", after: " + after)
}

type Subject struct {
	context   string
	observers []*Observer
}

func (subject *Subject) addObserver(observer *Observer) {
	subject.observers = append(subject.observers, observer)
}

func (subject *Subject) notify() {
	for _, observer := range subject.observers {
		(*observer).Update(subject)
	}
}

func (subject *Subject) updateSubject(context string) {
	subject.context = context
	subject.notify()
}
