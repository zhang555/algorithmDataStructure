package sort1

type Queue struct {
	s []interface{}
}

func (q *Queue) Push(x interface{}) {
	q.s = append(q.s, x)
}

func (q *Queue) Pop() interface{} {
	if q.Empty() {
		return 0
	}

	v := q.s[0]
	q.s = q.s[1:]
	return v
}

func (q *Queue) Empty() bool {
	return len(q.s) == 0
}

func (q *Queue) Len() int {
	return len(q.s)
}

func (q *Queue) Top() interface{} {
	if q.Empty() {
		return 0
	}
	return q.s[len(q.s)-1]
}
