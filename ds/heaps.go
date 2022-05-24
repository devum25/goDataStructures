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
	for idx > 1 && h.compare(idx){
         h.swap(Parent(idx),idx)
		 idx = Parent(idx)
	}
}
// Pop

// Top

// compare
func (h *Heap) compare(idx int) bool{
    if h.Type == MIN{
        return h.Array[Parent(idx)] > h.Array[idx]
	}else{
       return h.Array[Parent(idx)] < h.Array[idx]
	}
}

