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
	vertex string
	weight float64
}

//拓扑排序
//先拿到所有顶点的入度 ， 如果是0 就入队 ，
//循环将入度为0的顶点出队 ，每出队一个，就把对应的顶点的入度减一
func Sort(g Graph) []string {
	vertexToIndegree := make(map[string]int)

	for vertex := range g.Entity {
		vertexToIndegree[vertex] = 0
	}

	for _, edges := range g.Entity {
		for _, edge := range edges {
			vertexToIndegree[edge.vertex]++
		}
	}

	queue := Queue{}

	for vertex := range g.Entity {
		v, ok := vertexToIndegree[vertex]
		if !ok {
			panic(` ! ok `)
		}
		if v == 0 {
			queue.Push(vertex)
		}
	}

	var result []string

	for !queue.Empty() {

		vertex := queue.Pop().(string)
		result = append(result, vertex)

		for _, edge := range g.Entity[vertex] {

			vertexToIndegree[edge.vertex]--

			v, ok := vertexToIndegree[edge.vertex]
			if !ok {
				panic(` ! ok `)
			}
			if v == 0 {
				queue.Push(edge.vertex)
			}
		}

	}

	return result
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
