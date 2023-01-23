package main

import (
	"fmt"

	"github.com/daniial79/data-structures-GO/dataStructures/graph"
)

func main() {
	mg := graph.GenGraph()

	mg.AddVertex("A")
	mg.AddVertex("B")
	mg.AddVertex("C")
	mg.AddVertex("D")

	// fmt.Printf("%+v\n", mg)

	mg.AddEdge("A", "B")
	mg.AddEdge("A", "C")
	mg.AddEdge("A", "D")

	mg.AddEdge("C", "B")
	mg.AddEdge("C", "D")

	fmt.Printf("%+v\n", mg)

	mg.RemoveVertex("B")

	fmt.Printf("%+v\n", mg)

}
