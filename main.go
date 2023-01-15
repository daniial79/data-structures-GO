package main

import (
	dll "github.com/daniial79/data-structures-GO/dataStructures/doublyLinkedList"
)

func main() {
	mdll := dll.GenDoublyLinkedList()

	mdll.Preppend(10)
	mdll.PrintList()

	mdll.Preppend(9)
	mdll.PrintList()

}
