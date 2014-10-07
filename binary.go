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

func (o Or) Eval() bool {
	res := false
	for _, n := range o.nodes {
		res = res || n.Eval()
	}
	return res
}

func (a And) Eval() bool {
	res := true
	for _, n := range a.nodes {
		res = res && n.Eval()
	}
	return res
}

func (x Xor) Eval() bool {
	res := true
	for _, n := range x.nodes {
		res = (res || n.Eval()) && !(n.Eval() && res)
	}
	return res
}
