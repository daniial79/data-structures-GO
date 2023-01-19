package bst

type node struct {
	value int
	right *node
	left  *node
}

type Bst struct {
	Root *node
}

func GenBst() *Bst {
	return &Bst{}
}
