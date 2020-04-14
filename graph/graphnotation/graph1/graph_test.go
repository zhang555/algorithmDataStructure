package shortestpath

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/gyuho/goraph"
)

func TestName(t *testing.T) {

	g := Graph{}

	g.Node = append(g.Node,
		[]int{1, 1, 1},
		[]int{1, 1, 1},
		[]int{1, 1, 1},
	)

	graph := goraph.NewGraph()

	for index, e := range g.Node {

		nd1, err := graph.GetNode(goraph.StringID(strconv.Itoa(index)))
		if err != nil {
			nd1 = goraph.NewNode(strconv.Itoa(index))
			graph.AddNode(nd1)

		}

		for index2, v := range e {
			nd2, err := graph.GetNode(goraph.StringID(strconv.Itoa(index2)))
			if err != nil {
				nd2 = goraph.NewNode(strconv.Itoa(index2))
				graph.AddNode(nd2)
			}
			graph.ReplaceEdge(nd1.ID(), nd2.ID(), float64(v))
		}

	}

	fmt.Println(graph.String())

}
