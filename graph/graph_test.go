package graph

import (
	"testing"

	. "github.com/minaorangina/data-structures-go/utils"
)

func TestDijkstra(t *testing.T) {
	graph := initGraph()

	weight, path := dijkstra(graph, "a", "i", "abcdefghi")

	AssertEqual(t, weight, 8)
	AssertEqual(t, path, "a c d g i")
}

func initGraph() graph {
	mappings := map[string]*node{}
	vals := "abcdefghi"
	g := graph{nodes: map[string][]edge{}}

	for _, c := range vals {
		mappings[string(c)] = &node{value: string(c)}
	}

	g.nodes["a"] = []edge{{mappings["b"], 5}, {mappings["c"], 3}, {mappings["e"], 2}}
	g.nodes["b"] = []edge{{mappings["d"], 2}}
	g.nodes["c"] = []edge{{mappings["b"], 1}, {mappings["d"], 1}}
	g.nodes["d"] = []edge{{mappings["a"], 1}, {mappings["g"], 2}}
	g.nodes["e"] = []edge{{mappings["a"], 1}, {mappings["e"], 7}, {mappings["h"], 4}}
	g.nodes["f"] = []edge{{mappings["b"], 3}, {mappings["f"], 1}}
	g.nodes["g"] = []edge{{mappings["c"], 3}, {mappings["i"], 2}}
	g.nodes["h"] = []edge{{mappings["c"], 2}, {mappings["f"], 2}, {mappings["g"], 2}}
	g.nodes["i"] = []edge{}

	return g
}
