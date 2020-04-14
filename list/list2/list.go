package list

type List struct {
	Head *Node
	Tail *Node
	Size int
}

type Node struct {
	Val  interface{}
	Next *Node
	Prev *Node
}

func New() *List {

	l := List{}

	head := &Node{}
	tail := &Node{Prev: head}
	head.Next = tail

	return &l
}

func (l *List) PushBack(node *Node) {
	node.Prev = l.Tail.Prev
	node.Next = l.Tail

	node.Prev.Next = node
	node.Next.Prev = node
	l.Size++
}

func (l *List) Remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	l.Size--
}
