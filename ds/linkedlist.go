package ds

import (
	"errors"
	"fmt"
)

func main(){
	l := &LinkedList{}
	l.Insert(0)
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Insert(4)

   l.Print()
  //1 2 3 4 5
 
  comp := make(map[int]*Node)
    
  ptr := l.head
  
  count := 0
  
  
  for ptr != nil{
	  
	  if ptr.next != nil{
	  comp[ptr.val] = ptr.next
	  }else{
		  comp[ptr.val] = nil
	  }
		   
	  ptr = ptr.next
  }
  
  nums := make([]int,0)
  nums = append(nums, 0,3,1,4)

  x := make(map[int]bool)
  
  for i := 0;i < len(nums);i++{
	  
	  for j := 0;j < len(nums);j++{
		  if comp[nums[i]] != nil{
				  if comp[nums[i]].val == nums[j]{
					  x[nums[j]] = true
					  count++
				  }
		  }else{
			  y := x[nums[j]]
			  if y == false{
			  count++
			  break
			  }
		  }
	  }
  }
}

type Node struct{
	val int
	next *Node
}

type LinkedList struct{
	head *Node
	length int
}

func New() *LinkedList{
	return &LinkedList{}
}
// insert new node at the end of the linked list
func (l *LinkedList) Insert(val int){
	n := Node{}
	n.val = val
	if l.length == 0{
		l.head = &n
		l.length++
		return
	}

	ptr := l.head

	for i:=0;i<l.length;i++{
         if ptr.next == nil{
			 ptr.next = &n
			 l.length++
			 return
		 }
		 ptr = ptr.next
	}
}

// InsertAt inserts new node at given position
func (l *LinkedList) InsertAt(pos int, value int) {
	// create a new node
	newNode := Node{}
	newNode.val = value
	// validate the position
	if pos < 0 {
		return
	}
	if pos == 0 {
		l.head = &newNode
		l.length++
		return
	}
	if pos > l.length {
		return
	}
	n := l.GetAt(pos)
	newNode.next = n
	prevNode := l.GetAt(pos - 1)
	prevNode.next = &newNode
	l.length++
}

// GetAt returns node at given position from linked list
func (l *LinkedList) GetAt(pos int) *Node {
	ptr := l.head
	if pos < 0 {
		return ptr
	}
	if pos > (l.length - 1) {
		return nil
	}
	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	return ptr
}

// Print displays all the nodes from linked list
func (l *LinkedList) Print() {
	if l.length == 0 {
		fmt.Println("No nodes in list")
	}
	ptr := l.head
	for i := 0; i < l.length; i++ {
		fmt.Println("Node: ", ptr)
		ptr = ptr.next
	}
}

//searching a node from linkedlist
func (l *LinkedList) Search(val int) int{
	ptr := l.head

	for i:=0;i<l.length;i++{
		if ptr.val == val{
			return i
		}
		ptr = ptr.next
	}
	return -1
}

//DeleteAt deletes node at given position from linkedlist
func (l *LinkedList) DeleteAt(pos int) error{
	// validate the position
	if pos < 0{
		fmt.Println("position can not be negative")
		return errors.New("position can not be negative")
	}
	if l.length == 0{
		fmt.Println("No nodes in list")
		return errors.New("No nodes in list")
	}
   prevNode := l.GetAt(pos -1)
   if prevNode == nil{
	   fmt.Println("Node not found")
	   return errors.New("Node not found")
   }

   prevNode.next = l.GetAt(pos).next
   l.length--
   return nil
}

// DeleteVal deletes node having given value from linked list
func (l *LinkedList) DeleteVal(val int) error {
	ptr := l.head
	if l.length == 0 {
		fmt.Println("List is empty")
		return errors.New("List is empty")
	}
	for i := 0; i < l.length; i++ {
		if ptr.val == val {
			if i > 0 {
				prevNode := l.GetAt(i - 1)
				prevNode.next = l.GetAt(i).next
			} else {
				l.head = ptr.next
			}
			l.length--
			return nil
		}
		ptr = ptr.next
	}
	fmt.Println("Node not found")
	return errors.New("Node not found")
}