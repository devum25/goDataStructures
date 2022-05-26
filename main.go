package main

import (
	"fmt"

	"github.com/dev/datastructures/ds"
)

func main(){
      table := ds.NewHashTable()

      table.Insert("Devum","Developer")
      table.Delete("Devum")
      fmt.Println(table.Search("Devum"))
   }
