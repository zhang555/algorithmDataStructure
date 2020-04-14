package short1

import (
	"fmt"
	"log"
	"testing"

	"github.com/gyuho/goraph"
)

type Graph struct {
	Entity map[string][]Edge
}

type Edge struct {
	vertex string
}

type ForShortest struct {
	Index    int
	Previous string
}

//无权重的情况下，找到最短路径
func ShortestPath(g Graph, start string, end string) ([]string, error) {
	vertexKnown := make(map[string]ForShortest)
	if _, ok := g.Entity[start]; !ok {
		return nil, nil
	}
	if _, ok := g.Entity[end]; !ok {
		return nil, nil
	}
	vertexKnown[start] = ForShortest{Index: 0}
	for i := 0; i < len(g.Entity); i++ {
		log.Println(vertexKnown)
		for vertex, shortest := range vertexKnown {
			if shortest.Index == i {
				for _, edge := range g.Entity[vertex] {
					_, ok := vertexKnown[edge.vertex]
					if ok {
						continue
					}
					forShortest, ok := vertexKnown[vertex]
					if !ok {
						panic(` !ok `)
					}
					vertexKnown[edge.vertex] = ForShortest{Index: forShortest.Index + 1, Previous: vertex}
				}
			}
		}
	}

	var result []string
	n := end
	result = append([]string{n}, result...)

	for {

		if n == start {
			break
		}

		forShortest, ok := vertexKnown[n]
		if !ok {
			return nil, fmt.Errorf(` dont have path `)
		}
		n = forShortest.Previous
		result = append([]string{n}, result...)

	}
	return result, nil
}

func TestName(t *testing.T) {

	g1 := Graph{}

	g1.Entity = make(map[string][]Edge)

	g1.Entity[`1`] = []Edge{
		{`2`},
		{`4`},
	}

	g1.Entity[`2`] = []Edge{
		{`4`},
		{`5`},
	}

	g1.Entity[`3`] = []Edge{
		{`1`},
		{`6`},
	}

	g1.Entity[`4`] = []Edge{
		{`3`},
		{`5`},
		{`6`},
		{`7`},
	}

	g1.Entity[`5`] = []Edge{
		{`7`},
	}

	g1.Entity[`6`] = nil

	g1.Entity[`7`] = []Edge{
		{`6`},
	}

	result, err := ShortestPath(g1, `3`, `7`)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)

	g := goraph.NewGraph()
	for key, mm := range g1.Entity {
		nd1, err := g.GetNode(goraph.StringID(key))
		if err != nil {
			nd1 = goraph.NewNode(key)
			g.AddNode(nd1)
		}
		for _, node := range mm {
			nd2, err := g.GetNode(goraph.StringID(node.vertex))
			if err != nil {
				nd2 = goraph.NewNode(node.vertex)
				g.AddNode(nd2)
			}
			g.ReplaceEdge(nd1.ID(), nd2.ID(), 1)
		}
	}

	ids, _, err := goraph.Dijkstra(g, goraph.StringID(`3`), goraph.StringID(`7`))
	if err != nil {
		log.Fatal(err)

	}
	log.Println(ids)

}
