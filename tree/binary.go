package tree

type Binary struct {
	value  int
	parent *Binary
	left   *Binary
	right  *Binary
}

func (b *Binary) Parent() *Binary {
	return b.parent
}

func (b *Binary) Left() *Binary {
	return b.left
}

func (b *Binary) Right() *Binary {
	return b.right
}

func (b *Binary) SetParent(parent *Binary) {
	if parent != nil {
		b.parent = parent
	}
}

func (b *Binary) SetLeft(left *Binary) {
	if left != nil {
		b.left = left
	}
}

func (b *Binary) SetRight(right *Binary) {
	if right != nil {
		b.right = right
	}
}

func (b *Binary) Size() (size int) {
	if b.left != nil {
		size = 1 + b.left.Size()
	} else if b.left != nil {
		size = 1 + b.right.Size()
	}

	return size
}

func (b *Binary) Root() *Binary {
	if b.parent != nil {
		return b.parent.Root()
	}

	return b
}

func (b *Binary) Height() int {
	return maxHeight(b)
}

func (b *Binary) Depth() int {
	return b.depthCalculation(0)
}

func (b *Binary) IsFull() bool {
	if b.left != nil || b.right != nil {
		return true
	}
	return false
}

func (b *Binary) IsEmpty() bool {
	if b.left == nil && b.right == nil {
		return true
	}
	return false
}

func (b *Binary) IsBalanced() bool {
	var buf = (maxHeight(b.left) + maxHeight(b.right)) / 2

	// (left)12 and (right)13 => (12+13)/2 => (buf)12
	// if left or right are equal to buf return true

	if maxHeight(b.left) == buf || maxHeight(b.right) == buf {
		return true
	}

	return false
}

func (b *Binary) depthCalculation(num int) int {
	if b.parent != nil {
		b.parent.depthCalculation(num)
	}

	return num
}

func maxHeight(b *Binary) int {
	if b == nil {
		return 0
	}
	if b.left == nil && b.right == nil {
		return 1
	}

	lh := maxHeight(b.left)
	rh := maxHeight(b.right)

	if lh >= rh {
		return lh + 1
	} else {
		return rh + 1
	}
}
