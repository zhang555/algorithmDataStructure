package shortestpath

type GraphNode struct {
	Vertex   int // 顶点
	Edge     int // 边
	Known    bool
	Distance int
}

//github.com/gyuho/goraph

func main() {

	//goraph.ID()
	//goraph.NewGraph()
}

//var Graph00 = goraph.Graph{
//	Name:           "graph_00",
//	TotalNodeCount: 8,
//	TotalEdgeCount: 30,
//	IsDAG:          false,
//	EdgeToWeight: []EdgeToWeight{
//		{[]string{"S", "A"}, 100.0},
//		{[]string{"S", "B"}, 14.0},
//		{[]string{"S", "C"}, 200.0},
//
//		{[]string{"A", "S"}, 15.0},
//		{[]string{"A", "B"}, 5.0},
//		{[]string{"A", "D"}, 20.0},
//		{[]string{"A", "T"}, 44.0},
//
//		{[]string{"B", "S"}, 14.0},
//		{[]string{"B", "A"}, 5.0},
//		{[]string{"B", "D"}, 30.0},
//		{[]string{"B", "E"}, 18.0},
//
//		{[]string{"C", "S"}, 9.0},
//		{[]string{"C", "E"}, 24.0},
//
//		{[]string{"D", "A"}, 20.0},
//		{[]string{"D", "B"}, 30.0},
//		{[]string{"D", "E"}, 2.0},
//		{[]string{"D", "F"}, 11.0},
//		{[]string{"D", "T"}, 16.0},
//
//		{[]string{"E", "B"}, 18.0},
//		{[]string{"E", "C"}, 24.0},
//		{[]string{"E", "D"}, 2.0},
//		{[]string{"E", "F"}, 6.0},
//		{[]string{"E", "T"}, 19.0},
//
//		{[]string{"F", "D"}, 11.0},
//		{[]string{"F", "E"}, 6.0},
//		{[]string{"F", "T"}, 6.0},
//
//		{[]string{"T", "A"}, 44.0},
//		{[]string{"T", "D"}, 16.0},
//		{[]string{"T", "F"}, 6.0},
//		{[]string{"T", "E"}, 19.0},
//	},
//}
