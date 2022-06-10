package main

import (
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
	a int
}

var singleInstance *single

func getInstance1(group *sync.WaitGroup) *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			//fmt.Println("Creating single instance now.")
			singleInstance = &single{1}
		} else {
			//fmt.Println("Single instance already created.")
		}
	} else {
		//fmt.Println("Single instance already created.")
	}

	group.Done()
	return singleInstance
}

func clear1() {
	singleInstance = nil
}
