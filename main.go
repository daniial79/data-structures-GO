package main

import dataStructure "github.com/daniial79/data-structures-GO/dataStructures"

func main() {
	mll := dataStructure.GenLinkedList()

	mll.PrintList()

	for i := 1; i <= 10; i++ {
		mll.Append(i)
	}

	mll.PrintList()
}
