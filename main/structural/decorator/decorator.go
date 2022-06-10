package main

type Drink interface {
	getPrice() int
}

type Espresso struct {
}

func (*Espresso) getPrice() int {
	return 10
}

type Sugar struct {
	drink Drink
}

func (sugar *Sugar) getPrice() int {
	return sugar.drink.getPrice() + 7
}

type Milk struct {
	drink Drink
}

func (milk *Milk) getPrice() int {
	return milk.drink.getPrice() + 10
}
