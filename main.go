package main

import (
	"fmt"

	dataStructure "github.com/daniial79/data-structures-GO/dataStructures"
)

func main() {
	mll := dataStructure.GenLinkedList()

	mll.Append(1)
	fmt.Printf("%+v\n", mll.PopFirst())

	for i := 1; i <= 10; i++ {
		mll.Append(i)
	}

	mll.PrintList()

	fmt.Printf("%+v\n", mll.PopFirst())
	mll.PrintList()

}
