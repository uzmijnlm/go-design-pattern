package main

import (
	"fmt"
	"log"
)

func main() {
	thread := &thread{}

	initializedState := &InitializedState{
		thread: thread,
	}
	runningState := &RunningState{
		thread: thread,
	}
	suspendedState := &SuspendedState{
		thread: thread,
	}
	succeededState := &SucceedState{
		thread: thread,
	}

	thread.initialized = initializedState
	thread.running = runningState
	thread.suspended = suspendedState
	thread.succeeded = succeededState
	thread.currentState = initializedState

	fmt.Printf("%T\n", thread.currentState)

	err := thread.run()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%T\n", thread.currentState)

	err = thread.suspend()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%T\n", thread.currentState)

	err = thread.run()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%T\n", thread.currentState)

	//err = thread.run()
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//fmt.Printf("%T\n", thread.currentState)

	err = thread.success()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%T\n", thread.currentState)
}
