package main

import (
	"fmt"

	"github.com/dev/datastructures/ds"
)

func main(){
      table := ds.NewHashTable()

      table.Insert("Devum","Developer")
      fmt.Println(table.Search("Devum"))
   }
