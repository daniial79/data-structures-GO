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

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (g *Graph) RemoveEdge(vertex1, vertex2 string) bool {
	if g.AdjList[vertex1] == nil || g.AdjList[vertex2] == nil {
		return false
	}

	var (
		founded bool
		index   int
	)

	for i, val := range g.AdjList[vertex1] {
		if val == vertex2 {
			founded = true
			index = i
		}
	}

	if founded {
		g.AdjList[vertex1] = remove(g.AdjList[vertex1], index)
	}

	for i, val := range g.AdjList[vertex2] {
		if val == vertex1 {
			founded = true
			index = i
		}
	}

	if founded {
		g.AdjList[vertex2] = remove(g.AdjList[vertex2], index)
	}

	return true
}

func (g *Graph) RemoveVertex(vertex string) bool {
	if _, founded := g.AdjList[vertex]; !founded {
		return false
	}

	for _, ver := range g.AdjList[vertex] {
		g.RemoveEdge(ver, vertex)
	}

	delete(g.AdjList, vertex)

	return true
}
