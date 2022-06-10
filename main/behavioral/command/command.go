package main

type Command interface {
	Execute()
	Undo()
}

type Operation struct {
	result int
}

func (op *Operation) addNum(num int) {
	op.result = op.result + num
}

func (op *Operation) subNum(num int) {
	op.result = op.result - num
}

type AddCommand struct {
	op  *Operation
	num int
}

type SubCommand struct {
	op  *Operation
	num int
}

func (command *AddCommand) Execute() {
	command.op.addNum(command.num)
}

func (command *AddCommand) Undo() {
	command.op.subNum(command.num)
}

func (command *SubCommand) Execute() {
	command.op.subNum(command.num)
}

func (command *SubCommand) Undo() {
	command.op.addNum(command.num)
}
