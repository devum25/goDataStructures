package ds

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int { return len(h)}
func (h IntHeap) Less(i,j int) bool { return h[i] < h[j]}
func (h IntHeap) Swap(i,j int) {h[i],h[j] = h[j],h[i]}

func (h *IntHeap) Push(x interface{}){
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{}{
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h IntHeap) Top() int{
	return h[0]
}

func TestHeap(){
	h := &IntHeap{50,40,70,60}
	heap.Init(h)
	heap.Push(h,100)
	heap.Push(h,110)

	//fmt.Print(h.Top())
  
	// fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}

}