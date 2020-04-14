package shortestpath

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph"
)

func TestName(t *testing.T) {

	g := Graph1{}

	g.Edges = append(g.Edges,
		Edge{`a`, `b`, 1},
		Edge{`a`, `c`, 1},
		Edge{`b`, `c`, 1},
	)

	graph := goraph.NewGraph()

	for _, e := range g.Edges {

		n1, err := graph.GetNode(goraph.StringID(e.n1))
		if err != nil {
			n1 = goraph.NewNode(e.n1)
			graph.AddNode(n1)

		}

		n2, err := graph.GetNode(goraph.StringID(e.n2))
		if err != nil {
			n2 = goraph.NewNode(e.n2)
			graph.AddNode(n2)
		}

		graph.ReplaceEdge(n1.ID(), n2.ID(), e.v)

	}

	fmt.Println(graph.String())

}
