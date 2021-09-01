package main


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
