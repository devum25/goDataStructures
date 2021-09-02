package main

import (
	"fmt"

	"github.com/dev/datastructures/ds"
)

func main(){
   l := ds.New()
   l.Insert(5)
   l.Insert(9)
   l.Insert(10)
   l.Insert(7)
   l.Insert(80)
   l.InsertAt(3,11)

   l.Print()
   fmt.Println("Post delete")

   l.DeleteAt(4)
   
   l.Print()
}