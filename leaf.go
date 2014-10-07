package main

type Leaf struct {
}

func (l Leaf) Eval(attr *Attributes) bool {
	return true
}
