package main

import (
	"fmt"

	stack "github.com/daniial79/data-structures-GO/dataStructures/stack"
)

func main() {
	ms := stack.GenStack()

	node, err := ms.Pop()
	fmt.Println(err)
	_ = node

	for i := 0; i <= 9; i++ {
		ms.Push(i)
	}

	ms.PrintStack()

}
