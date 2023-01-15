package main

import (
	dll "github.com/daniial79/data-structures-GO/dataStructures/doublyLinkedList"
)

func main() {
	mdll := dll.GenDoublyLinkedList()

	// mdll.PrintList()

	mdll.Append(1)
	mdll.Append(2)
	mdll.Append(3)
	mdll.Append(4)
	mdll.Append(5)

	mdll.PrintList()
}
