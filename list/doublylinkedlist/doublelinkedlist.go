package doublylinkedlist

type List struct {
	First *Element
	Last  *Element
	size  uint
}

type Element struct {
	Value interface{}
	prev  *Element
	next  *Element
}

func New(value ...interface{}) *List {
	l := &List{}
	for _, v := range value {
		l.PushBack(v)
	}
	return l
}

func (l *List) PushFront(value interface{}) *Element {
	if l.size == 0 {
		e := &Element{Value: value}
		l.First = e
		l.Last = e
		l.size++
		return e
	} else {
		e := &Element{Value: value}
		oldFirst := l.First
		e.next = oldFirst
		oldFirst.prev = e
		l.First = e
		l.size++
		return e
	}
}

func (l *List) PushBack(value interface{}) *Element {
	if l.size == 0 {
		e := &Element{Value: value}
		l.First = e
		l.Last = e
		l.size++
		return e
	} else {
		e := &Element{Value: value}
		last := l.Last
		l.Last.next = e
		e.prev = last
		l.Last = e
		l.size++
		return e
	}
}

//将 e1 移动到 e2 后面
func (l *List) Move(e1 *Element, e2 *Element) {
	if e1 == nil || e2 == nil {
		return
	}

	if e1 == e2 {
		return
	}

	if e1.prev == e2 {
		return
	}

	switch {
	case e1 == l.First && e2 == l.Last:
		l.First = e1.next
		l.Last = e1

	case e1 == l.First:
		l.First = e1.next

	case e2 == l.Last:
		l.Last = e1

	case e1 == l.Last:
		l.Last = e1.prev
	}

	if e1.next != nil {
		e1.next.prev = e1.prev
	}

	if e1.prev != nil {
		e1.prev.next = e1.next
	}

	e3 := e2.next
	e2.next = e1
	e1.prev = e2

	if e3 != nil {
		e3.prev = e1
	}
	e1.next = e3

}

func (l *List) MoveToFirst(e1 *Element) {
	if e1 == nil {
		return
	}

	if l.size >= 2 {
		l.Move(e1, l.First)
		l.Move(e1.prev, e1)
	}
}

func (l *List) MoveToLast(e1 *Element) {
	if e1 == nil {
		return
	}

	l.Move(e1, l.Last)
}

func (l *List) Remove(e1 *Element) {
	if e1 == nil {
		return
	}

	switch {
	case e1 == l.First:
		l.First = e1.next
	case e1 == l.Last:
		l.Last = e1.prev
	}

	if e1.next != nil {
		e1.next.prev = e1.prev
	}

	if e1.prev != nil {
		e1.prev.next = e1.next
	}
	l.size--
}

func (e Element) Next() *Element {
	return e.next
}

func (e Element) Prev() *Element {
	return e.prev
}
