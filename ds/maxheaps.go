package ds

// MaxHeap struct has a slice that holds the array
type MaxHeap struct{
	Array []int
}

func NewMaxHeap() *MaxHeap{
  return  &MaxHeap{Array: []int{-1}}
}

// Insert adds an element to that heap
func (m *MaxHeap) Insert(key int){
	m.Array = append(m.Array, key)

	m.maxHeapifyUp(len(m.Array)-1)
}



// Extract return the largest key, and removes it from heap


func (m *MaxHeap) maxHeapifyUp(index int){
     for index>1 && m.Array[parent(index)] < m.Array[index]{
		 m.swap(parent(index),index)
	     index = parent(index)
	 }
}


func parent(index int) int{
     return index/2
}

func rightChild(index int) int{
	return (2*index+1)
}

func leftChild(index int) int{
	return (2*index)
}

func (m *MaxHeap) swap(i1,i2 int){
    m.Array[i1],m.Array[i2] = m.Array[i2],m.Array[i1]
}

func (m *MaxHeap) Top() int{
	return m.Array[1]
}

func (m *MaxHeap) Pop() int{
     idx := len(m.Array)-1
	 m.swap(1,idx)
	 popped := m.Array[idx]
	 m.Array = m.Array[:idx]
     m.heapify(1)

	 return popped
}

func (m *MaxHeap) heapify(i int){
   l := leftChild(i)
   r := rightChild(i)


   maxIdx := i

   if l < len(m.Array) && m.Array[i] < m.Array[l]{
	   maxIdx = l
   }

   if (r < len(m.Array)) && (m.Array[maxIdx] < m.Array[r]){
	   maxIdx = r
   }

   if maxIdx != i{
	   m.swap(maxIdx,i)
	   m.heapify(maxIdx)
   }
}

