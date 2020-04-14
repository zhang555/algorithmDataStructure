package lfu21

type List2 struct {
	Head *Node
	Tail *Node
	Size int
}

type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

func New() *List2 {

	l := List2{}

	head := &Node{}
	tail := &Node{Prev: head}
	head.Next = tail

	l.Head = head
	l.Tail = tail

	return &l
}

func (l *List2) PushBack(node *Node) {
	node.Prev = l.Tail.Prev
	node.Next = l.Tail

	node.Prev.Next = node
	node.Next.Prev = node
	l.Size++
}

func (l *List2) PushFront(node *Node) {
	node.Next = l.Head.Next
	node.Prev = l.Head

	node.Prev.Next = node
	node.Next.Prev = node
	l.Size++
}

func (l *List2) Remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	l.Size--
}

func (l *List2) Back() *Node {
	if l.Size == 0 {
		return nil
	}
	return l.Tail.Prev
}

func (l *List2) Front() *Node {
	if l.Size == 0 {
		return nil
	}
	return l.Head.Next
}

func (l *List2) Len() int {
	return l.Size
}
