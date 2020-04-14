package shortestpath

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph"
)

func TestName(t *testing.T) {

	g1 := Graph{}

	g1.Entity = make(map[string][]Edge)

	g1.Entity[`a`] = []Edge{
		{`b`, 1},
		{`c`, 1},
	}

	g1.Entity[`b`] = []Edge{
		{`c`, 1},
		{`d`, 1},
	}

	g1.Entity[`c`] = nil

	g1.Entity[`d`] = []Edge{
		{`a`, 1},
	}

	g := goraph.NewGraph()
	for key, mm := range g1.Entity {
		nd1, err := g.GetNode(goraph.StringID(key))
		if err != nil {
			nd1 = goraph.NewNode(key)
			g.AddNode(nd1)
		}
		for _, node := range mm {
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
