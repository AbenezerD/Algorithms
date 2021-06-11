package week2

import (
	"container/heap"
)

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {return h[i] > h[j]}

func (h maxHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	maxH := maxHeap(stones)
	h := &maxH
	heap.Init(h)
	for ; h.Len() > 0;  {
		maxStone := heap.Pop(h)
		if h.Len() == 0 {
			return maxStone.(int)
		}
		secondMax := heap.Pop(h)
		diff := maxStone.(int) - secondMax.(int)
		if diff > 0 {
			heap.Push(h, diff)
		}
	}

	return 0
}

//				 	1
//			2				1
//     7		8		4


//				 	8
//			7				4
//     1		2		1

