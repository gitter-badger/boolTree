package main

import "fmt"

type Leaf struct {
	Key      string
	Value    string
	EvalFunc func(l Leaf, attr *Attributes) bool
	Name     string
}

func (l Leaf) Eval(attr *Attributes) bool {
	return l.EvalFunc(l, attr)
}

func (l Leaf) Print() string {
	return "(" + fmt.Sprintf("f(%s) == %s", l.Key, l.Value) + ")"
}

func IsIn(l Leaf, attr *Attributes) bool {
	if val, ok := attr.Values[l.Key]; ok {
		for _, v := range val {
			if v == l.Value {
				return true
			}
		}
	}
	return false
}
