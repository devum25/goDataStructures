package main

import (
	"fmt"

	"github.com/dev/datastructures/ds"
)

func main(){
   // // l := ds.New()
   // // l.Insert(5)
   // // l.Insert(9)
   // // l.Insert(10)
   // // l.Insert(7)
   // // l.Insert(80)
   // // l.InsertAt(3,11)

   // // l.Print()
   // // fmt.Println("Post delete")

   // // l.DeleteAt(4)
   
   // // l.Print()
   // l := New()
   // l.Insert(1)
   // l.Insert(2)
   // l.Insert(3)
   // l.Insert(4)

   // // l.DeleteAt(2)
   // // l.Search(2)
   // l.GetAt(1)

   // heap := ds.NewHeap()

   // heap.Array = append(heap.Array,100,40,50,20,30,15)

   // heap.Insert(120)

   // fmt.Println(heap.Array)

   // fmt.Println(heap.Pop())

   // fmt.Println(heap.Array)


   heap := ds.NewMinHeap()

   heap.Array = append(heap.Array, 10,30,40,50,55,65,60)
   
   heap.Insert(5)

   fmt.Println(heap.Array)

   for i:=0;i<5;i++{

   
   fmt.Println(heap.Pop())

   fmt.Println(heap.Array)
   }
}