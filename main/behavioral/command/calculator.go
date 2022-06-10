package main

type Calculator struct {
	commands     []Command
	undoCommands []Command
}

func (calculator *Calculator) addCommand(command Command) {
	calculator.commands = append(calculator.commands, command)
}

func (calculator *Calculator) press() {
	commandToDo := calculator.commands[0]
	calculator.commands = calculator.commands[1:]
	commandToDo.Execute()
	calculator.undoCommands = append(calculator.undoCommands, commandToDo)
}

func (calculator *Calculator) undo() {
	commandToUndo := calculator.undoCommands[len(calculator.undoCommands)-1]
	calculator.undoCommands = calculator.undoCommands[:len(calculator.undoCommands)-1]
	commandToUndo.Undo()
	calculator.commands = append([]Command{commandToUndo}, calculator.commands...)
}
