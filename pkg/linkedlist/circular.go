package linkedlist

type nodeCircular struct {
	data string
	next *nodeCircular
}

type Circular struct {
	Head *nodeCircular
	Last *nodeCircular
	len  int
}


func (ll *Circular) Add(d string) {
	node := &nodeCircular{data: d}
	ll.len++
	
	if ll.Head == nil {
		ll.Head, ll.Last = node, node

		return
	}

	node.next, ll.Head = ll.Head, node
}

func (ll *Circular) Remove(d string) (ok bool) {
	if ll.Head.data == d {
		ll.Head = ll.Head.next
		ll.len--

		return true
	}

	return ll.removeFirst(ll.Head.next, d)
}

func (ll *Circular) Contains(d string) (ok bool) {
	if ll.Head.data == d {
		return true
	}

	return ll.findFirst(ll.Head.next, d)
}

func (ll *Circular) Clear() {
	ll.Head = nil
	ll.Last = nil
	ll.len = 0
}

func (ll *Circular) CopyToSlice() []string {
	arr  := make([]string, ll.len)
	prev := ll.Head
	pres := prev.next

	for i := 0; i < ll.len; i++ {
		arr = append(arr, prev.data)
		prev, pres = pres, pres.next
	}

	return arr
}

func (ll *Circular) Len() int {
	return ll.len
}

func (ll *Circular) removeFirst(n *nodeCircular, d string) (ok bool) {
	if ll.Last == n {
		return false
	} else if n.next.data == d {
		n.next = n.next.next
		ll.len--

		return true
	}

	return ll.removeFirst(n.next, d)
}

func (ll *Circular) findFirst(n *nodeCircular, d string) (ok bool) {
	if ll.Head == n {
		return false
	} else if n.data == d {
		return true
	}

	return ll.findFirst(n.next, d)
}

