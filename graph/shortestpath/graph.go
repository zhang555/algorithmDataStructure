package shortestpath

//弧表示法
type Graph1 struct {
	Edges []Edge
}

type Edge struct {
	n1 string
	n2 string
	v  float64
}

//邻接表 slice
type Graph2 struct {
	Edges map[string][]Edge
}
