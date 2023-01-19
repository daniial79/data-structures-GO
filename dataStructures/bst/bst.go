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

func (b *Bst) Contains(val int) bool {
	if b.Root == nil {
		return false
	}

	temp := b.Root
	for temp != nil {
		if temp.value == val {
			return true
		} else if temp.value > val {
			temp = temp.left
		} else if temp.value < val {
			temp = temp.right
		}
	}

	return false
}

func (b *Bst) Bfs() []int {
	if b.Root == nil {
		return nil
	}

	currentNode := b.Root
	queue, result := []*node{}, []int{}

	queue = append(queue, currentNode)

	for len(queue) != 0 {
		currentNode = queue[0]
		queue = queue[1:]

		result = append(result, currentNode.value)

		if currentNode.left != nil {
			queue = append(queue, currentNode.left)
		}

		if currentNode.right != nil {
			queue = append(queue, currentNode.right)
		}
	}

	return result
}

func (b *Bst) PreOrderDfs() []int {
	if b.Root == nil {
		return nil
	}

	result := []int{}

	var traverse func(currentNode *node)
	traverse = func(currentNode *node) {
		result = append(result, currentNode.value)

		if currentNode.left != nil {
			traverse(currentNode.left)
		}

		if currentNode.right != nil {
			traverse(currentNode.right)
		}
	}

	traverse(b.Root)

	return result
}

func (b *Bst) InOrderDfs() []int {
	if b.Root == nil {
		return nil
	}
	result := []int{}

	var traverse func(currentNode *node)
	traverse = func(currentNode *node) {

		if currentNode.left != nil {
			traverse(currentNode.left)
		}

		result = append(result, currentNode.value)

		if currentNode.right != nil {
			traverse(currentNode.right)
		}
	}

	traverse(b.Root)

	return result

}

func (b *Bst) PostOrderDfs() []int {
	if b.Root == nil {
		return nil
	}
	result := []int{}

	var traverse func(currentNode *node)
	traverse = func(currentNode *node) {

		if currentNode.left != nil {
			traverse(currentNode.left)
		}

		if currentNode.right != nil {
			traverse(currentNode.right)
		}

		result = append(result, currentNode.value)

	}

	traverse(b.Root)

	return result

}
