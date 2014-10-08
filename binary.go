package main

import "strings"

type Binary struct {
	Nodes []Operation
}

type Or struct {
	Binary
}

type And struct {
	Binary
}

type Xor struct {
	Binary
}

func (o Or) Eval(attr *Attributes) bool {
	res := false
	for _, n := range o.Nodes {
		res = res || n.Eval(attr)
	}
	return res
}

func (o Or) Print() string {
	res := make([]string, len(o.Nodes))
	for i, n := range o.Nodes {
		res[i] = n.Print()
	}
	resString := "(" + strings.Join(res, " || ") + ")"
	return resString
}

func (a And) Eval(attr *Attributes) bool {
	res := true
	for _, n := range a.Nodes {
		res = res && n.Eval(attr)
	}
	return res
}

func (a And) Print() string {
	res := make([]string, len(a.Nodes))
	for i, n := range a.Nodes {
		res[i] = n.Print()
	}
	resString := "(" + strings.Join(res, " && ") + ")"
	return resString
}

func (x Xor) Eval(attr *Attributes) bool {
	res := true
	for _, n := range x.Nodes {
		res = (res || n.Eval(attr)) && !(n.Eval(attr) && res)
	}
	return res
}

func (x Xor) Print() string {
	res := make([]string, len(x.Nodes))
	for i, n := range x.Nodes {
		res[i] = n.Print()
	}
	resString := "(" + strings.Join(res, " XOR ") + ")"
	return resString
}
