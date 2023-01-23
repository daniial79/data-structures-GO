package graph

type Graph struct {
	AdjList map[string][]string
}

func GenGraph() *Graph {
	return &Graph{
		AdjList: make(map[string][]string),
	}
}

func (g *Graph) AddVertex(vertex string) bool {
	if _, found := g.AdjList[vertex]; found {
		return false
	}

	g.AdjList[vertex] = make([]string, 0)

	return true
}
