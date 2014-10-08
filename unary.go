package main

import "fmt"

type Unary struct {
	node Operation
}

func (u Unary) Node(n Operation) {
	u.node = n
}

type Not struct {
	Unary
}

func (n Not) Eval(attr *Attributes) bool {
	return !n.node.Eval(attr)
}

func (n Not) Print() string {
	return fmt.Sprintf("!%s", n.node.Print())
}
