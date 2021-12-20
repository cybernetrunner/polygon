package adt

type nodeQueue struct {
	data string
	prev *nodeQueue
	next *nodeQueue
}

type Queue struct {
	Head *nodeQueue
	Tail *nodeQueue
}

func (q Queue) Pop() *nodeQueue {
	node := q.Tail
	q.Tail = q.Tail.prev

	return node
}

func (q Queue) Push(data string) {
	node := &nodeQueue{
		data: data,
		next: q.Head,
	}
	q.Head = node
}
