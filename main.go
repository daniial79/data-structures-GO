package main

import (
	"fmt"

	"github.com/daniial79/data-structures-GO/dataStructures/graph"
)

func main() {
	mg := graph.GenGraph()

	mg.AddVertex("A")

	mg.AddVertex("B")

	fmt.Printf("%+v\n", mg)

	mg.AddEdge("A", "B")

	fmt.Printf("%+v\n", mg)

}
