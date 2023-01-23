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

func (g *Graph) AddEdge(vertex1, vertex2 string) bool {
	if g.AdjList[vertex1] == nil || g.AdjList[vertex2] == nil {
		return false
	}

	for _, ver := range g.AdjList[vertex1] {
		if ver == vertex2 {
			return false
		}
	}

	g.AdjList[vertex1] = append(g.AdjList[vertex1], vertex2)
	g.AdjList[vertex2] = append(g.AdjList[vertex2], vertex1)

	return true
}
