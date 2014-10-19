package priority_queue

import (
	"container/heap"
)

type Interface interface {
	Less(other interface{}) bool
}

type sorter []Interface

func (s *sorter) Push(x interface{}) {
	*s = append(*s, x.(Interface))
}

func (s *sorter) Pop() interface{} {
	n := len(*s)
	if n > 0 {
		x := (*s)[n-1]
		*s = (*s)[0 : n-1]
		return x
	}
	return nil
}

func (s *sorter) Top() interface{} {
	if len(*s) > 0 {
		return (*s)[0]
	}
	return nil
}

func (s *sorter) Len() int {
	return len(*s)
}

func (s *sorter) Less(i, j int) bool {
	return (*s)[i].Less((*s)[j])
}

func (s *sorter) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

type PriorityQueue struct {
	s *sorter
}

func New() *PriorityQueue {
	q := &PriorityQueue{s: new(sorter)}
	heap.Init(q.s)
	return q
}

func (q *PriorityQueue) Push(x Interface) {
	heap.Push(q.s, x)
}

func (q *PriorityQueue) Pop() Interface {
	return heap.Pop(q.s).(Interface)
}

func (q *PriorityQueue) Top() Interface {
	return q.s.Top().(Interface)
}

func (q *PriorityQueue) Len() int {
	return q.s.Len()
}
