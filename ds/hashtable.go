package ds


const ArraySize = 7

//Hashtable structure
// Hashtable will hold an array of type bucket
type HashTable struct{
	array [ArraySize]*bucket
}

func NewHashTable() *HashTable{
        h := &HashTable{}
		for i:=0;i<len(h.array);i++{
			h.array[i] = &bucket{}
		}
		return h
}

//bucket structure
// bucket is a linkedlist in each slot of the hashtable
type bucket struct{
	head *bucketNode
}

// bucketNode structure
type bucketNode struct{
	key string
	val interface{}
	next *bucketNode
}
//Insert
func (h *HashTable) Insert(key string,val interface{}){
	index := hash(key)
	h.array[index].insert(key,val)
}

//Search
func (h *HashTable) Search(key string) interface{}{
	index := hash(key)
	return h.array[index].search(key)
	// return nil
}

//Delete
func (h *HashTable) Delete(key string){
	 index := hash(key)
	 h.array[index].delete(key)
}

// bucket insert
func (b *bucket) insert(k string,v interface{}){
	newNode := &bucketNode{key: k,val: v}
	newNode.next = b.head
	b.head = newNode
}
// bucket search
func (b *bucket) search(k string) interface{}{
	currentNode := b.head
	for currentNode != nil{
		if currentNode.key == k{
           return currentNode.val
		}
		currentNode = currentNode.next
	}

	return nil
}
// bucket delete
func (b *bucket) delete(k string){
   
	if b.head.key == k{
		b.head = b.head.next
		return
	}
	
	previousNode := b.head

	for previousNode.next != nil{
		if previousNode.next.key == k{
			previousNode.next = previousNode.next.next
			return
		}else{
			previousNode = previousNode.next
		}
	}
}

// hash function
func hash(key string) int{
	sum := 0
	for _,v := range key{
         sum += int(v)
	}

	return (sum%ArraySize)
}