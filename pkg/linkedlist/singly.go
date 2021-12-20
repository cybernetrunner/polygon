package linkedlist

type nodeSingly struct {
	data string
	next *nodeSingly
}

type Singly struct {
	Head *nodeSingly
	Tail *nodeSingly
	len  int
}

func (ll *Singly) Add(d string) {
	n := &nodeSingly{data: d, next: nil}

	ll.Tail = n
	ll.len++

	if ll.Head == nil {
		ll.Head = n

		return
	}

	ll.writeToLast(ll.Head, n)
}

func (ll *Singly) Remove(d string) (ok bool) {
	if ll.Head.data == d {
		ll.Head = ll.Head.next
		ll.len--

		return true
	}

	return ll.removeFirst(ll.Head.next, d)
}

func (ll *Singly) Contains(d string) (ok bool) {
	if ll.Head.data == d {
		return true
	}

	return ll.findFirst(ll.Head.next, d)
}

func (ll *Singly) Clear() {
	ll.Head = nil
	ll.Tail = nil
	ll.len = 0
}

func (ll *Singly) CopyToSlice() []string {
	arr  := make([]string, ll.len)
	prev := ll.Head
	pres := prev.next

	for i := 0; i < ll.len; i++ {
		arr = append(arr, prev.data)
		prev, pres = pres, pres.next
	}

	return arr
}

func (ll *Singly) Len() int {
	return ll.len
}

func (ll *Singly) writeToLast(last, added *nodeSingly) {
	if last.next != nil {
		ll.writeToLast(last.next, added)
	}

	last.next = added
}

func (ll *Singly) removeFirst(n *nodeSingly, d string) (ok bool) {
	if ll.Tail == n {
		return false
	} else if n.next.data == d {
		n.next = n.next.next
		ll.len--

		return true
	}

	return ll.removeFirst(n.next, d)
}

func (ll *Singly) findFirst(n *nodeSingly, d string) (ok bool) {
	if n.data == d {
		return true
	}

	return ll.findFirst(n.next, d)
}
