package main

import "fmt"

type thread struct {
	initialized state
	running     state
	suspended   state
	succeeded   state

	currentState state
}

func (thread *thread) run() error {
	return thread.currentState.run()
}

func (thread *thread) suspend() error {
	return thread.currentState.suspend()
}

func (thread *thread) success() error {
	return thread.currentState.success()
}

type state interface {
	run() error
	suspend() error
	success() error
}

type InitializedState struct {
	thread *thread
}

func (state *InitializedState) run() error {
	state.thread.currentState = state.thread.running
	return nil
}

func (state *InitializedState) suspend() error {
	return fmt.Errorf("cannot from initialized to suspend")
}

func (state *InitializedState) success() error {
	return fmt.Errorf("cannot from initialized to success")
}

type RunningState struct {
	thread *thread
}

func (state *RunningState) run() error {
	return fmt.Errorf("cannot from running to running")
}

func (state *RunningState) suspend() error {
	state.thread.currentState = state.thread.suspended
	return nil
}

func (state *RunningState) success() error {
	state.thread.currentState = state.thread.succeeded
	return nil
}

type SuspendedState struct {
	thread *thread
}

func (state *SuspendedState) run() error {
	state.thread.currentState = state.thread.running
	return nil
}

func (state *SuspendedState) suspend() error {
	return fmt.Errorf("cannot from suspended to suspended")
}

func (state *SuspendedState) success() error {
	return fmt.Errorf("cannot from suspended to success")
}

type SucceedState struct {
	thread *thread
}

func (state *SucceedState) run() error {
	return fmt.Errorf("cannot from success to running")
}

func (state *SucceedState) suspend() error {
	return fmt.Errorf("cannot from success to suspended")
}

func (state *SucceedState) success() error {
	return fmt.Errorf("cannot from success to success")
}
