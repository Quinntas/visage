package queue

type Queue []interface{}

func (q *Queue) Push(x interface{}) {
	*q = append(*q, x)
}

func (q *Queue) Pop() interface{} {
	var el interface{}
	el, *q = (*q)[0], (*q)[1:]
	return el
}

func (q *Queue) Len() int {
	return len(*q)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Clear() {
	*q = (*q)[0:0]
}

func NewQueue() *Queue {
	return &Queue{}
}
