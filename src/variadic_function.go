package main

import (
  "fmt"
)

func  main()  {
  variadicFunction("10",11,12,13,14,15,16)
}

func variadicFunction(sequence string, numbers ...int)  {
  for _, number := range numbers {
    fmt.Println(number)
  }

}
