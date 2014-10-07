package main

type Constant interface {
	Eval() bool
}

type True struct {
}

func (t True) Eval() bool {
	return true
}

type False struct {
}

func (f False) Eval() bool {
	return false
}
