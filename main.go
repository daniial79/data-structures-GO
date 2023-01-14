package main

import (
	"fmt"

	dataStructure "github.com/daniial79/data-structures-GO/dataStructures"
)

func main() {
	mll := dataStructure.GenLinkedList()

	for i := 1; i <= 10; i++ {
		mll.Append(i)
	}

	mll.PrintList()

	fmt.Println(mll.Get(5))
}
