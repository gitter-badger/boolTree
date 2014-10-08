package main

import "fmt"

type Unary struct {
	Node Operation
}

type Not struct {
	Unary
}

func (n Not) Eval(attr *Attributes) bool {
	return !n.Node.Eval(attr)
}

func (n Not) Print() string {
	return fmt.Sprintf("!%s", n.Node.Print())
}
