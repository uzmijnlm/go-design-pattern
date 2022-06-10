package main

type Row struct {
	cols map[string]int
}

func (row *Row) getByColName(colName string) int {
	col, ok := row.cols[colName]
	if ok {
		return col
	} else {
		return 0
	}
}

type Expression interface {
	eval(*Row) int
}

type ColumnExpression struct {
	field string
}

func (expr *ColumnExpression) eval(row *Row) int {
	return row.getByColName(expr.field)
}

type ConstantExpression struct {
	value int
}

func (expr *ConstantExpression) eval(row *Row) int {
	return expr.value
}

type AddExpression struct {
	left  *Expression
	right *Expression
}

func (expr *AddExpression) eval(row *Row) int {
	return (*expr.left).eval(row) + (*expr.right).eval(row)
}
