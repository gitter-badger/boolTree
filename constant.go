package main

type Constant interface {
	Eval() bool
}

type True struct {
}

func (t True) Eval(attr *Attributes) bool {
	return true
}

type False struct {
}

func (f False) Eval(attr *Attributes) bool {
	return false
}
