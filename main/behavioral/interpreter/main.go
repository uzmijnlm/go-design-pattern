package main

import "fmt"

func main() {
	row := &Row{cols: make(map[string]int)}
	row.cols["id"] = 0

	var left, right, add Expression
	left = &ColumnExpression{
		field: "id",
	}
	right = &ConstantExpression{
		value: 1,
	}
	add = &AddExpression{
		left:  &left,
		right: &right,
	}

	fmt.Println(add.eval(row))
}
