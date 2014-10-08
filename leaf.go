package main

type Leaf struct {
	Key   string
	Value string
}

func (l Leaf) Eval(attr *Attributes) bool {
	if val, ok := attr.Values[l.Key]; ok {
		for _, v := range val {
			if v == l.Value {
				return true
			}
		}
	}
	return false
}

func (l Leaf) Print() string {
	return ""
}
