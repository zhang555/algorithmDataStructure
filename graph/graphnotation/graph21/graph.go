package shortestpath

//邻接表 map
type Graph struct {
	Entity map[string][]Edge
}

type Edge struct {
	n string
	v float64
}
