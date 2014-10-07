package main

type Binary struct {
	nodes []Operation
}

func (b Binary) Nodes(ns []Operation) {
	b.nodes = append(b.nodes, ns...)
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
	for _, n := range o.nodes {
		res = res || n.Eval(attr)
	}
	return res
}

func (a And) Eval(attr *Attributes) bool {
	res := true
	for _, n := range a.nodes {
		res = res && n.Eval(attr)
	}
	return res
}

func (x Xor) Eval(attr *Attributes) bool {
	res := true
	for _, n := range x.nodes {
		res = (res || n.Eval(attr)) && !(n.Eval(attr) && res)
	}
	return res
}
