package ds

type MinHeap struct {
	Array []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{Array: []int{-1}}
}


// insert
func (h *MinHeap) Insert(data int){
	h.Array = append(h.Array, data)
   
	 h.minHeapUp(len(h.Array)-1)
	
}

// parent
func parentindex(i int) int{
	return (i/2)
}

//left child
func left(i int) int{
	return (2*i)
}

// right child
func right(i int) int{
	return (2*i + 1)
}

// swap
func (h *MinHeap) swap(i1,i2 int){
    h.Array[i1],h.Array[i2] = h.Array[i2],h.Array[i1]
}

// top
func (h *MinHeap) Top() int{
	return h.Array[1]
}

// pop
func (h *MinHeap) Pop() int{
	idx := len(h.Array)-1
	h.swap(1,idx)
	popped := h.Array[idx]

	h.Array = h.Array[:idx]

	h.heapify(1)

	return popped
}

// heapify
func (h *MinHeap) heapify(idx int){
	l := left(idx)
	r := right(idx)

	minIdx := idx

	if l <= (len(h.Array)-1) && h.Array[minIdx] > h.Array[l]{
		minIdx = l
	}

	if r <= (len(h.Array)-1) && h.Array[minIdx] > h.Array[r]{
		minIdx = r
	}

	if minIdx != idx{
		h.swap(minIdx,idx)
		h.heapify(minIdx)
	}
}

// minHeapUp
func (h *MinHeap) minHeapUp(idx int){
    for idx > 1 && h.Array[parentindex(idx)] > h.Array[idx]{
         h.swap(parentindex(idx),idx)
		 idx = parentindex(idx)
	}
} 