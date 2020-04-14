package graph3

//弧表示法
type Graph struct {
	Edges []Edge
}

type Edge struct {
	n1 string
	n2 string
	v  float64
}
