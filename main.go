package main

import (
	"fmt"

	bst "github.com/daniial79/data-structures-GO/dataStructures/bst"
)

func main() {
	mbst := bst.GenBst()

	mbst.Insert(50)
	mbst.Insert(75)
	mbst.Insert(25)
	mbst.Insert(80)
	mbst.Insert(70)
	mbst.Insert(30)
	mbst.Insert(20)

	fmt.Println(mbst.Bfs())
	fmt.Println(mbst.PreOrderDfs())
	fmt.Println(mbst.InOrderDfs())
	fmt.Println(mbst.PostOrderDfs())

}
