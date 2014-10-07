package main

type Operation interface {
	Eval(attr *Attributes) bool
}
