package main

type Unary struct {
	node Operation
}

func (u Unary) Node(n Operation) {
	u.node = n
}

type Not struct {
	Unary
}

func (n Not) Eval() bool {
	return !n.node.Eval()
}
