package stack

type node struct {
	Value int
	Rear  *node
}

type Stack struct {
	Top    *node
	Height int
}
