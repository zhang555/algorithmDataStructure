package doublylinkedlist

import (
	"log"
	"testing"
)

func isSameValueList(l1 List, l2 List) bool {
	return isSameValueListByFirst(l1.First, l2.First) && isSameValueListByLast(l1.Last, l2.Last) && l1.size == l2.size
}

func isSameValueListByFirst(l1 *Element, l2 *Element) bool {
	switch {
	case l1 == nil && l2 == nil:
		return true
	case l1 == nil:
		return false

	case l2 == nil:
		return false

	default:
		if l1.Value.(int) != l2.Value.(int) {
			return false
		}
		return isSameValueListByFirst(l1.Next(), l2.Next())
	}
}

func isSameValueListByLast(l1 *Element, l2 *Element) bool {
	switch {
	case l1 == nil && l2 == nil:
		return true
	case l1 == nil:
		return false

	case l2 == nil:
		return false

	default:
		if l1.Value.(int) != l2.Value.(int) {
			return false
		}
		return isSameValueListByLast(l1.Prev(), l2.Prev())
	}
}

func TestIsSameValueList1(t *testing.T) {

	var (
		e1 *Element = &Element{Value: 1}
		e2 *Element = &Element{Value: 2}
		e3 *Element = &Element{Value: 3}

		l1 List = List{}
		l2 List = List{}
	)

	e1.next = e2

	e2.prev = e1
	e2.next = e3

	e3.prev = e2
	l1.First = e1
	l1.Last = e3
	l1.size = 3

	l2.PushBack(1)
	l2.PushBack(2)
	l2.PushBack(3)

	if !isSameValueList(l1, l2) {
		t.Error()
	}

	l2 = List{}
	l2.PushFront(3)
	l2.PushFront(2)
	l2.PushFront(1)

	if !isSameValueList(l1, l2) {
		t.Error()
	}

}

func TestList_Move1(t *testing.T) {

	var (
		e1 *Element = &Element{Value: 1}
		e2 *Element = &Element{Value: 2}
		e3 *Element = &Element{Value: 3}

		l1 List = List{}
		l2 List = List{}
	)

	e1.next = e2

	e2.prev = e1
	e2.next = e3

	e3.prev = e2
	l1.First = e1
	l1.Last = e3
	l1.size = 3

	l1.Move(e1, e2)

	l2.PushBack(2)
	l2.PushBack(1)
	l2.PushBack(3)

	if !isSameValueList(l1, l2) {
		t.Error(` should 2 1 3 `)
	}

}

func TestList_Move2(t *testing.T) {

	var (
		e1 *Element = &Element{Value: 1}

		l1 List = List{}
		l2 List = List{}
	)

	l1.First = e1
	l1.Last = e1
	l1.size = 1

	l1.Move(e1, e1)

	l2.PushBack(1)

	if !isSameValueList(l1, l2) {
		t.Error()
	}

}

func TestList_Move3(t *testing.T) {

	var (
		e1 *Element = &Element{Value: 1}
		e2 *Element = &Element{Value: 2}

		l1 List = List{}
		l2 List = List{}
	)

	e1.next = e2

	e2.prev = e1

	l1.First = e1
	l1.Last = e2
	l1.size = 2

	l1.Move(e1, e2)

	l2.PushBack(2)
	l2.PushBack(1)

	if !isSameValueList(l1, l2) {
		t.Error()
	}

}

func TestList_Move4(t *testing.T) {

	var (
		e1 *Element = &Element{Value: 1}
		e2 *Element = &Element{Value: 2}
		e3 *Element = &Element{Value: 3}

		l1 List = List{}
		l2 List = List{}
	)

	e1.next = e2

	e2.prev = e1
	e2.next = e3

	e3.prev = e2
	l1.First = e1
	l1.Last = e3
	l1.size = 3

	l1.Move(e3, e1)

	l2.PushBack(1)
	l2.PushBack(3)
	l2.PushBack(2)

	if !isSameValueList(l1, l2) {
		t.Error()
	}

}

func TestList_Move5(t *testing.T) {

	var (
		e1 *Element = &Element{Value: 1}
		e2 *Element = &Element{Value: 2}
		e3 *Element = &Element{Value: 3}

		l1 List = List{}
		l2 List = List{}
	)

	e1.next = e2

	e2.prev = e1
	e2.next = e3

	e3.prev = e2
	l1.First = e1
	l1.Last = e3
	l1.size = 3

	l1.Move(e3, e2)

	l2.PushBack(1)
	l2.PushBack(2)
	l2.PushBack(3)

	if !isSameValueList(l1, l2) {
		t.Error()
	}

}

func TestList_Move6(t *testing.T) {

	var (
		e1 *Element = &Element{Value: 1}
		e2 *Element = &Element{Value: 2}
		e3 *Element = &Element{Value: 3}

		l1 List = List{}
		l2 List = List{}
	)

	e1.next = e2

	e2.prev = e1
	e2.next = e3

	e3.prev = e2
	l1.First = e1
	l1.Last = e3
	l1.size = 3

	l1.Move(e2, e3)

	l2.PushBack(1)
	l2.PushBack(3)
	l2.PushBack(2)

	if !isSameValueList(l1, l2) {
		t.Error()
	}

}

func TestList_Move7(t *testing.T) {
	var (
		e1 *Element = &Element{Value: 1}
		e2 *Element = &Element{Value: 2}
		e3 *Element = &Element{Value: 3}

		l1 List = List{}
		l2 List = List{}
	)

	e1.next = e2

	e2.prev = e1
	e2.next = e3

	e3.prev = e2
	l1.First = e1
	l1.Last = e3
	l1.size = 3

	l1.Move(e1, e3)

	l2.PushBack(2)
	l2.PushBack(3)
	l2.PushBack(1)

	if !isSameValueList(l1, l2) {
		t.Error()
	}

}

func TestList_MoveToFirst1(t *testing.T) {
	var (
		e1 *Element = &Element{Value: 1}
		e2 *Element = &Element{Value: 2}
		e3 *Element = &Element{Value: 3}

		l1 List = List{}
		l2 List = List{}
	)

	e1.next = e2

	e2.prev = e1
	e2.next = e3

	e3.prev = e2
	l1.First = e1
	l1.Last = e3
	l1.size = 3

	l1.MoveToFirst(e3)

	l2.PushBack(3)
	l2.PushBack(1)
	l2.PushBack(2)

	if !isSameValueList(l1, l2) {
		t.Error()
	}

}

func TestList_MoveToFirst2(t *testing.T) {
	var (
		e1 *Element = &Element{Value: 1}
		e2 *Element = &Element{Value: 2}
		e3 *Element = &Element{Value: 3}

		l1 List = List{}
		l2 List = List{}
	)

	e1.next = e2

	e2.prev = e1
	e2.next = e3

	e3.prev = e2
	l1.First = e1
	l1.Last = e3
	l1.size = 3

	l1.MoveToFirst(e2)

	l2.PushBack(2)
	l2.PushBack(1)
	l2.PushBack(3)

	if !isSameValueList(l1, l2) {
		t.Error()
	}

}

func TestList_MoveToFirst3(t *testing.T) {
	var (
		e1 *Element = &Element{Value: 1}

		l1 List = List{}
		l2 List = List{}
	)

	l1.First = e1
	l1.Last = e1
	l1.size = 1

	l1.MoveToFirst(e1)

	l2.PushBack(1)

	if !isSameValueList(l1, l2) {
		t.Error()
	}

}

func TestList_MoveToLast1(t *testing.T) {
	var (
		e1 *Element = &Element{Value: 1}
		e2 *Element = &Element{Value: 2}
		e3 *Element = &Element{Value: 3}

		l1 List = List{}
		l2 List = List{}
	)

	e1.next = e2

	e2.prev = e1
	e2.next = e3

	e3.prev = e2
	l1.First = e1
	l1.Last = e3
	l1.size = 3

	l1.MoveToLast(e1)

	l2.PushBack(2)
	l2.PushBack(3)
	l2.PushBack(1)

	if !isSameValueList(l1, l2) {
		t.Error()
	}

}

func TestList_MoveToLast2(t *testing.T) {
	var (
		e1 *Element = &Element{Value: 1}
		e2 *Element = &Element{Value: 2}
		e3 *Element = &Element{Value: 3}

		l1 List = List{}
		l2 List = List{}
	)

	e1.next = e2

	e2.prev = e1
	e2.next = e3

	e3.prev = e2
	l1.First = e1
	l1.Last = e3
	l1.size = 3

	l1.MoveToLast(e2)

	l2.PushBack(1)
	l2.PushBack(3)
	l2.PushBack(2)

	if !isSameValueList(l1, l2) {
		t.Error()
	}

}

func TestList_MoveToLast3(t *testing.T) {
	var (
		e1 *Element = &Element{Value: 1}
		e2 *Element = &Element{Value: 2}
		e3 *Element = &Element{Value: 3}

		l1 List = List{}
		l2 List = List{}
	)

	e1.next = e2

	e2.prev = e1
	e2.next = e3

	e3.prev = e2
	l1.First = e1
	l1.Last = e3
	l1.size = 3

	l1.MoveToLast(e3)

	l2.PushBack(1)
	l2.PushBack(2)
	l2.PushBack(3)

	if !isSameValueList(l1, l2) {
		t.Error()
	}

}

func TestList_Remove1(t *testing.T) {
	var (
		e1 *Element = &Element{Value: 1}
		e2 *Element = &Element{Value: 2}
		e3 *Element = &Element{Value: 3}

		l1 List = List{}
		l2 List = List{}
	)

	e1.next = e2

	e2.prev = e1
	e2.next = e3

	e3.prev = e2
	l1.First = e1
	l1.Last = e3
	l1.size = 3

	l1.Remove(e1)

	l2.PushBack(2)
	l2.PushBack(3)

	if !isSameValueList(l1, l2) {
		t.Error()
	}

}

func TestList_Remove2(t *testing.T) {
	var (
		e1 *Element = &Element{Value: 1}
		e2 *Element = &Element{Value: 2}
		e3 *Element = &Element{Value: 3}

		l1 List = List{}
		l2 List = List{}
	)

	e1.next = e2

	e2.prev = e1
	e2.next = e3

	e3.prev = e2
	l1.First = e1
	l1.Last = e3
	l1.size = 3

	l1.Remove(e2)

	l2.PushBack(1)
	l2.PushBack(3)

	if !isSameValueList(l1, l2) {
		t.Error()
	}

}

func TestList_Remove3(t *testing.T) {
	var (
		e1 *Element = &Element{Value: 1}
		e2 *Element = &Element{Value: 2}
		e3 *Element = &Element{Value: 3}

		l1 List = List{}
		l2 List = List{}
	)

	e1.next = e2

	e2.prev = e1
	e2.next = e3

	e3.prev = e2
	l1.First = e1
	l1.Last = e3
	l1.size = 3

	l1.Remove(e3)

	l2.PushBack(1)
	l2.PushBack(2)

	if !isSameValueList(l1, l2) {
		t.Error()
	}

}

func TestList_Remove13(t *testing.T) {
	var (
		e1 *Element = &Element{Value: 1}
		e2 *Element = &Element{Value: 2}
		e3 *Element = &Element{Value: 3}
		e4 *Element = &Element{Value: 4}

		l1 List = List{}
		l2 List = List{}
	)

	e1.next = e2

	e2.prev = e1
	e2.next = e3

	e3.prev = e2
	l1.First = e1
	l1.Last = e3
	l1.size = 3

	l1.Remove(e4)

	l2.PushBack(1)
	l2.PushBack(2)
	l2.PushBack(3)

	if !isSameValueList(l1, l2) {
		t.Error()
	}

}

func TestList1(t *testing.T) {

}

func TestList_Iterator1(t *testing.T) {
	var (
		l1 List = List{}
	)

	l1.PushBack(1)
	l1.PushBack(2)
	l1.PushBack(3)

	iterator := l1.Iterator()
	for {
		v, ok := iterator.Next()
		if !ok {
			break
		}
		log.Println(v)
	}
}

func TestList_Iterator11(t *testing.T) {
	var (
		l1 List = List{}
	)

	l1.PushBack(1)
	l1.PushBack(2)
	l1.PushBack(3)

	iterator := l1.Iterator()

	l1.First.Value = 10
	l1.First.next.Value = 11
	l1.First.next.next.Value = 12
	for {
		v, ok := iterator.Next()
		if !ok {
			break
		}
		log.Println(v)
	}
}

func TestList_Iterator2(t *testing.T) {
	var (
		l1 List = List{}
	)

	l1.PushBack(1)

	iterator := l1.Iterator()
	for {
		v, ok := iterator.Next()
		if !ok {
			break
		}
		log.Println(v)
	}
}

func TestList_Iterator3(t *testing.T) {
	var (
		l1 List = List{}
	)

	iterator := l1.Iterator()

	for {
		v, ok := iterator.Next()
		if !ok {
			break
		}
		log.Println(v)
	}
}

func TestList_DeepIterator101(t *testing.T) {
	var (
		l1 List = List{}
	)

	l1.PushBack(1)
	l1.PushBack(2)
	l1.PushBack(3)

	iterator := l1.DeepIterator()

	l1.First.Value = 10
	l1.First.next.Value = 11
	l1.First.next.next.Value = 12

	for {
		v, ok := iterator.Next()
		if !ok {
			break
		}
		log.Println(v)
	}
}

func TestList_DeepIterator1011(t *testing.T) {
	var (
		l1 List = List{}
	)

	l1.PushBack(1)
	l1.PushBack(2)
	l1.PushBack(3)

	l1.First.Value = 10
	l1.First.next.Value = 11
	l1.First.next.next.Value = 12

	iterator := l1.DeepIterator()
	for {
		v, ok := iterator.Next()
		if !ok {
			break
		}
		log.Println(v)
	}
}

func TestList_DeepIterator102(t *testing.T) {
	var (
		l1 List = List{}
	)

	l1.PushBack(1)
	l1.PushBack(2)
	l1.PushBack(3)

	iterator := l1.DeepIterator()

	for {
		v, ok := iterator.Next()
		if !ok {
			break
		}
		log.Println(v)
	}
}

func TestList_DeepIterator103(t *testing.T) {
	var (
		l1 List = List{}
	)

	iterator := l1.DeepIterator()

	for {
		v, ok := iterator.Next()
		if !ok {
			break
		}
		log.Println(v)
	}
}

func TestList2001(t *testing.T) {
	var (
		l1 List = List{}
		m       = map[*Element]struct{}{}
	)

	//r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < 100; i++ {

		e := &Element{}
		m[e] = struct{}{}
	}

	for k := range m {
		l1.PushBack(k)
		l1.Remove(k)

	}

	//for k := range m {
	//}

}
