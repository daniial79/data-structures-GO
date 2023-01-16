package main

import (
	stack "github.com/daniial79/data-structures-GO/dataStructures/stack"
)

func main() {
	ms := stack.GenStack()

	for i := 0; i <= 9; i++ {
		ms.Push(i)
	}

	ms.PrintStack()

}
