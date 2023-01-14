package main

import (
	dataStructure "github.com/daniial79/data-structures-GO/dataStructures"
)

func main() {
	mll := dataStructure.GenLinkedList()

	for i := 1; i <= 10; i++ {
		mll.Append(i)
	}

	mll.PrintList()

	mll.SetValue(1000, 5)

	mll.PrintList()

}
