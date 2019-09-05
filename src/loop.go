package main

import (
  "fmt"
  // "time"
)

func main()  {
  // var i int
  // for i < 10 {
  //   fmt.Println(i)
  //   i = i + 1
  // }
  //
  // for j := 0; j < 10; j++ {
  //   fmt.Println(j)
  //   // time.Sleep(10 * time.Second)
  //   if j == 3{
  //     break
  //   }
  // }

  list_num := []int{1,2,3,4,5,6,7,8}
  for index, num := range list_num {
    fmt.Println(index, num)
  }

  for <expr> {
    code
    for <expr> {
      code
      for <expr>{
        code
        break breakPoint
      }
    }
  }


}
