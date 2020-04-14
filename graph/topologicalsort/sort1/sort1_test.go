package sort1

import (
	"log"
	"testing"
)

//邻接表 map
type Graph struct {
	Entity map[string][]Edge
}

type Edge struct {
	n string
	v float64
}

//拓扑排序， 先找到没有入边的点， 去掉这个点， 循环
//假设点的个数是V，时间复杂度  O(V2)
func Sort(g Graph) []string {
	var result []string
	l := len(g.Entity)
	for i := 0; i < l; i++ {
		r := findNewVertexOfIndegreeOfZero(g)
		if r == `` {
			return nil
		}
		result = append(result, r)
		delete(g.Entity, r)
	}
	return result
}

//找到没有入边的点
func findNewVertexOfIndegreeOfZero(g Graph) string {
	var edges []Edge
	for _, v := range g.Entity {
		edges = append(edges, v...)
	}
outer:
	for k := range g.Entity {
		for _, edge := range edges {
			if k == edge.n {
				continue outer
			}
		}
		return k
	}
	return ``
}

func TestName(t *testing.T) {

	g1 := Graph{}

	g1.Entity = make(map[string][]Edge)

	g1.Entity[`1`] = []Edge{
		{`2`, 1},
		{`3`, 1},
		{`4`, 1},
	}

	g1.Entity[`2`] = []Edge{
		{`4`, 1},
		{`5`, 1},
	}

	g1.Entity[`3`] = []Edge{
		{`6`, 1},
	}

	g1.Entity[`4`] = []Edge{
		{`3`, 1},
		{`6`, 1},
		{`7`, 1},
	}

	g1.Entity[`5`] = []Edge{
		{`4`, 1},
		{`7`, 1},
	}

	g1.Entity[`6`] = nil

	g1.Entity[`7`] = []Edge{
		{`6`, 1},
	}

	r := Sort(g1)
	log.Println(r)

}
