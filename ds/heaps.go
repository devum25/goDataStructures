package ds

const(
	MAX string = "Max"
	MIN string = "Min"
)

type Heap struct{
	Array []int
	Type string
}

func NewHeap(t string) *Heap{
	return &Heap{Array: []int{-1},Type: t}
}



func (h *Heap) Insert(data int){
     h.Array = append(h.Array, data)
	 h.HeapUp(len(h.Array)-1)
}

// Left
func Left(idx int) int{
	return (2*idx)
}

// Right
func Right(idx int) int{
	return (2*idx+1)
}

// Parent
func Parent(idx int) int{
	return (idx/2)
}

func (h *Heap) swap(i1,i2 int){
	h.Array[i1],h.Array[i2] = h.Array[i2],h.Array[i1]
}

// HeapUp
func (h *Heap) HeapUp(idx int){
	for idx > 1 && h.compare(Parent(idx),idx){
         h.swap(Parent(idx),idx)
		 idx = Parent(idx)
	}
}
// Pop
func (h *Heap) Pop() int{
	idx := len(h.Array)-1

	h.swap(1,idx)
	popped := h.Array[idx]

	h.Array = h.Array[:idx]
	h.heapify(1)

	return popped
}

// Top
func (h *Heap) Top() int{
	return h.Array[1]
}

// compare
func (h *Heap) compare(i1,i2 int) bool{
    if h.Type == MIN{
        return h.Array[i1] > h.Array[i2]
	}else{
       return h.Array[i1] < h.Array[i2]
	}
}

func (h *Heap) heapify(idx int){
	l := Left(idx)
	r := Right(idx)

	temp := idx

	if l <= len(h.Array)-1 && (h.compare(temp,l)){
		temp = l
	}

	if r <= len(h.Array)-1 && (h.compare(temp,r)){
         temp = r
	}

	if temp != idx{
		h.swap(temp,idx)
		h.heapify(temp)
	}
}