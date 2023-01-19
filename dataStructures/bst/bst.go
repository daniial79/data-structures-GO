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

func (b *Bst) Insert(val int) bool {
	node := node{value: val}
	if b.Root == nil {
		b.Root = &node
		return true
	}
	temp := b.Root

	for {
		if temp.value == node.value {
			break
		} else if temp.value > node.value {
			if temp.left == nil {
				temp.left = &node
				break
			}
			temp = temp.left
		} else if temp.value < node.value {
			if temp.right == nil {
				temp.right = &node
				break
			}
			temp = temp.right
		}
	}

	return true
}
