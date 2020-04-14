package shortestpath

//邻接表 slice
type Graph struct {
	Entity []Entity
}

type Entity struct {
	n     string
	edges []Edge
}

type Edge struct {
	n string
	v float64
}
