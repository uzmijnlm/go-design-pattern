package main

import "fmt"

func main() {
	operation := &Operation{
		result: 0,
	}

	command1 := &AddCommand{
		op:  operation,
		num: 2,
	}
	command2 := &SubCommand{
		op:  operation,
		num: 1,
	}
	command3 := &AddCommand{
		op:  operation,
		num: 3,
	}

	calculator := &Calculator{}
	calculator.addCommand(command1)
	calculator.addCommand(command2)
	calculator.addCommand(command3)

	calculator.press()
	fmt.Printf("after add 2, result is %d\n", operation.result)

	calculator.press()
	fmt.Printf("after sub 1, result is %d\n", operation.result)

	calculator.press()
	fmt.Printf("after add 3, result is %d\n", operation.result)

	calculator.undo()
	fmt.Printf("after undo add 3(which means sub 3), result is %d\n", operation.result)

	calculator.undo()
	fmt.Printf("after undo sub 1(which means add 1), result is %d\n", operation.result)
}
