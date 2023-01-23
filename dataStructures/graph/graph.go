package graph

type Graph struct {
	AdjList map[string][]string
}

func GenGraph() *Graph {
	return &Graph{}
}
