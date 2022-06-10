package main

type Operator interface {
	SetA(int)
	SetB(int)
	Operate() int
}

type OperatorFactory interface {
	Create() Operator
}

type PlusOperatorFactory struct {
}

type MinusOperatorFactory struct {
}

type BaseOperator struct {
	a, b int
}

func (operator *BaseOperator) SetA(a int) {
	operator.a = a
}

func (operator *BaseOperator) SetB(b int) {
	operator.b = b
}

type PlusOperator struct {
	int
	BaseOperator
}

type MinusOperator struct {
	BaseOperator
}

func (operator PlusOperator) Operate() int {
	return operator.a + operator.b
}

func (operator MinusOperator) Operate() int {
	return operator.a - operator.b
}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{}
}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{}
}
