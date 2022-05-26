package ds

import (
	"fmt"
	"strconv"
)


const ArraySize = 7
var count =0

//Hashtable structure
// Hashtable will hold an array of type bucket
type HashTable struct{
	array []*bucket
}

func NewHashTable(size int) *HashTable{
	length := 0
	if size == 0{
            length = ArraySize
	}else{
		length = size
	}
	   arr := make([]*bucket,length)
        h := &HashTable{array: arr}
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
	index := h.hash(key)
	h.array[index].insert(key,val)	
	count++
	load_factor := float32(count)/float32(len(h.array))
	if  load_factor > 0.7{
		h.rehash()
	}

}

//Search
func (h *HashTable) Search(key string) interface{}{
	index := h.hash(key)
	return h.array[index].search(key)
	// return nil
}

//Delete
func (h *HashTable) Delete(key string){
	 index := h.hash(key)
	 h.array[index].delete(key)
}

// bucket insert
func (b *bucket) insert(k string,v interface{}){
	newNode := &bucketNode{key: k,val: v}
	newNode.next = b.head
	b.head = newNode
}

func (h *HashTable) rehash(){
    // save the ptr to oldTable and we will do insertion in new table
	oldTable := *h

	count = 0

	// increase the table size 
	*h = *NewHashTable(2*ArraySize+1)
    
	// copy elements from old table to new table

	for _,v := range oldTable.array{
        currentNode := v.head
		for currentNode != nil{
			key := currentNode.key
            val := currentNode.val
			h.Insert(key,val)

			currentNode = currentNode.next
		}
	}

}

func (h *HashTable) Print(){
	for i,v := range h.array{
		currentNode := v.head
		fmt.Println()
		fmt.Print(strconv.Itoa(i))
		for currentNode != nil{
            fmt.Print("->"+currentNode.key)
			currentNode = currentNode.next
		}
	}
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
func (h *HashTable) hash(key string) int{
	sum := 0
	for _,v := range key{
         sum += int(v)
	}

	return (sum%len(h.array))
}