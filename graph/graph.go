package graph

import "math"

type node struct {
	value    string
	children []*node
}

type edge struct {
	neighbour *node
	weight    int
}

type graph struct {
	nodes map[string][]edge
}

func dijkstra(g graph, start, end, nodeNames string) (int, string) {
	pathWeight := map[string]int{}

	for _, c := range nodeNames {
		if string(c) == start {
			pathWeight[start] = 0
		} else {
			pathWeight[string(c)] = int(math.Inf(1))
		}
	}

	previous := map[string]*node{}
	_ = previous

	return 0, ""
}
