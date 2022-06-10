package main

import (
	"fmt"
)

type API interface {
	Say(name string) string
}

type hiAPI struct{}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

type helloAPI struct{}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func NewAPI(i int) API {
	if i == 0 {
		return &hiAPI{}
	} else if i == 1 {
		return &helloAPI{}
	}
	return nil
}
