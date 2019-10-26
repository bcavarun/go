// Arrays are single type fixed length
// Slicess are single type variable length. Slices are built on top of array
// Slices are references, always passed as refernece.


package main

import (
  "fmt"
)

func main()  {
  // Declares a slice called slice1
  // slice1 := make(<type>, <len>, <Cap>)
  //If cap is not specified, then cap is same as len.
  slice1 := make([]string, 5, 10)
  fmt.Println(len(slice1), cap(slice1))
}
