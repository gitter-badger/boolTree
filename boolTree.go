package main

import (
	"fmt"
	"log"
)

func main() {
	log.Print("Going on\n")

	attr := new(Attributes)
	attr.Values = make(map[string][]string)
	attr.Values["tag"] = make([]string, 2)
	attr.Values["tag"][0] = "news"
	attr.Values["tag"][1] = "mobile"

	left := Leaf{Key: "tag", Value: "news", EvalFunc: IsIn}
	right := Leaf{Key: "tag", Value: "mobile", EvalFunc: IsIn}
	andNode := And{}
	andNode.Nodes = []Operation{left, right}
	fmt.Printf("%s\n", andNode.Print())
	fmt.Printf("%v\n", andNode.Eval(attr))

	notNode := Not{}
	notNode.Node = andNode
	fmt.Printf("%s\n", notNode.Print())
	fmt.Printf("%v\n", notNode.Eval(attr))
}
