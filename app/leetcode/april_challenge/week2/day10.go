package week2

type MinStack struct {
	top *Node
}

type Node struct {
	min  int
	val  int
	next *Node
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if this.top == nil {
		this.top = &Node{min: x, val: x}
	} else {
		min := this.top.min
		if x < min {
			min = x
		}
		this.top = &Node{
			min:  min,
			val:  x,
			next: this.top,
		}
	}
}

func (this *MinStack) Pop() {
	if this.top != nil {
		this.top = this.top.next
	}
}

func (this *MinStack) Top() int {
	if this.top != nil {
		return this.top.val
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if this.top != nil {
		return this.top.min
	}

	return 0
}
