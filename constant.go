package main

type True struct {
}

func (t True) Eval(attr *Attributes) bool {
	return true
}

func (t True) Print() string {
	return "TRUE"
}

type False struct {
}

func (f False) Eval(attr *Attributes) bool {
	return false
}

func (f False) Print() string {
	return "FALSE"
}
