package adt

type nodeStack struct {
	data string
	next *nodeStack
}

type Stack struct {
	Head *nodeStack
}

func (s *Stack) Pop() *nodeStack {
	node := s.Head
	s.Head = s.Head.next

	return node
}

func (q *Stack) Push(data string) {
	q.Head = &nodeStack{
		data: data,
		next: q.Head,
	}
}
