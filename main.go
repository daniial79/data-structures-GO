package main

import (
	"fmt"

	dll "github.com/daniial79/data-structures-GO/dataStructures/doublyLinkedList"
)

func main() {
	mdll := dll.GenDoublyLinkedList()

	for i := 0; i <= 9; i++ {
		mdll.Append(i)
	}

	mdll.PrintList()

	fmt.Printf("%+v\n", mdll.Get(5))
}
