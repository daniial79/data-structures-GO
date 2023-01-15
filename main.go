package main

import (
	"fmt"

	dll "github.com/daniial79/data-structures-GO/dataStructures/doublyLinkedList"
)

func main() {
	mdll := dll.GenDoublyLinkedList()

	// mdll.PrintList()

	mdll.Append(-1)

	mdll.PrintList()

	fmt.Printf("%+v\n", mdll.Pop())

	for i := 0; i <= 9; i++ {
		mdll.Append(i)
	}

	mdll.PrintList()

	fmt.Printf("%+v\n", mdll.Pop())

	mdll.PrintList()
}
