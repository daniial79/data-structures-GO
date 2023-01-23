package main

import (
	"fmt"

	"github.com/daniial79/data-structures-GO/dataStructures/graph"
)

func main() {
	mg := graph.GenGraph()

	mg.AddVertex("A")

	mg.AddVertex("B")

	mg.AdjList["A"] = append(mg.AdjList["A"], "B")
	mg.AdjList["B"] = append(mg.AdjList["B"], "A")

	fmt.Printf("%+v\n", mg)

}
