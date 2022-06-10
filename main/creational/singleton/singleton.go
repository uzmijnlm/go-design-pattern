package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type singleton struct {
}

var instance *singleton

func getInstance(i int, group *sync.WaitGroup) *singleton {
	fmt.Println(i)
	if instance == nil {
		once.Do(func() {
			fmt.Println("Creating single instance now.")
			instance = &singleton{}
		})
	} else {
		fmt.Println("Single instance already created.")
	}
	group.Done()
	return instance
}

func clear() {
	instance = nil
}
