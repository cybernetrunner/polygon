package linkedlist

type nodeDoubly struct {
	data string
	prev *nodeDoubly
	next *nodeDoubly
}

type Doubly struct {
	Head *nodeDoubly
	Tail *nodeDoubly
	len  int
}

func (ll *Doubly) AddFirst(data string) {
	temp := ll.Head
	node := &nodeDoubly{
		data: data,
		next: temp,
	}

	temp.prev = node
	ll.Head   = node
}

func (ll *Doubly) AddLast(d string) {
	n := &nodeDoubly{data: d, next: nil}

	ll.Tail = n
	ll.len++

	if ll.Head == nil {
		ll.Head = n

		return
	}

	ll.writeToLast(ll.Head, n)
}

func (ll *Doubly) RemoveFirst(d string) (ok bool) {
	if ll.Head.data == d {
		ll.Head = ll.Head.next
		ll.len--

		return true
	}

	return ll.removeFirst(ll.Head.next, d)
}

func (ll *Doubly) RemoveLast(d string) (ok bool) {
	if ll.Tail.data == d {
		ll.Tail = ll.Tail.prev
		ll.len--

		return true
	}

	return ll.removeLast(ll.Tail.prev, d)
}

func (ll *Doubly) Contains(d string) (ok bool) {
	if ll.Head.data == d {
		return true
	}

	return ll.findFirst(ll.Head.next, d)
}

func (ll *Doubly) Clear() {
	ll.Head = nil
	ll.Tail = nil
	ll.len = 0
}

func (ll *Doubly) CopyToSlice() []string {
	arr  := make([]string, ll.len)
	prev := ll.Head
	pres := prev.next

	for i := 0; i < ll.len; i++ {
		arr = append(arr, prev.data)
		prev, pres = pres, pres.next
	}

	return arr
}

func (ll *Doubly) Len() int {
	return ll.len
}

func (ll *Doubly) writeToLast(last, added *nodeDoubly) {
	if last.next != nil {
		ll.writeToLast(last.next, added)
	}

	last.next = added
}

func (ll *Doubly) removeFirst(n *nodeDoubly, d string) (ok bool) {
	if ll.Tail == n {
		return false
	} else if n.next.data == d {
		n.next = n.next.next
		ll.len--

		return true
	}

	return ll.removeFirst(n.next, d)
}

func (ll *Doubly) removeLast(n *nodeDoubly, d string) (ok bool) {
	if ll.Tail == n {
		return false
	} else if n.prev.data == d {
		n.prev = n.prev.next
		ll.len--

		return true
	}

	return ll.removeLast(n.prev, d)
}

func (ll *Doubly) findFirst(n *nodeDoubly, d string) (ok bool) {
	if n.data == d {
		return true
	}

	return ll.findFirst(n.next, d)
}
