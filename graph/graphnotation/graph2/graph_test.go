package shortestpath

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph"
)

func TestName(t *testing.T) {

	g1 := Graph{}

	g1.Entity = append(g1.Entity,
		Entity{`a`, []Edge{
			{`b`, 1},
			{`c`, 1},
		}},

		Entity{`b`, []Edge{
			{`c`, 1},
			{`d`, 1},
		}},

		Entity{`c`, []Edge{}},

		Entity{`d`, []Edge{
			{`a`, 1},
		}},
	)

	g := goraph.NewGraph()
	for _, mm := range g1.Entity {
		nd1, err := g.GetNode(goraph.StringID(mm.n))
		if err != nil {
			nd1 = goraph.NewNode(mm.n)
			g.AddNode(nd1)
		}
		for _, node := range mm.edges {
			nd2, err := g.GetNode(goraph.StringID(node.n))
			if err != nil {
				nd2 = goraph.NewNode(node.n)
				g.AddNode(nd2)
			}
			g.ReplaceEdge(nd1.ID(), nd2.ID(), node.v)
		}
	}

	fmt.Println(g.String())

}
