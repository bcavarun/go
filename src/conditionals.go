package main

import (
  "fmt"
  "os"
)

func main()  {
  if false {
    fmt.Println("Success")
  } else if false {
    fmt.Println("In Elseif")
  } else {
    fmt.Println("Failure")
  }
  //
  // switch "Deep"  {
  //   // Not all the case below this will be executed like other languages. Implicit breaks in GO unlike other languages.
  //   case "Docker" : fmt.Println("Docker")
  //   case "Deep" : fmt.Println("Deep")
  //   case "Dive" : fmt.Println("Dive")
  //   default : fmt.Println("Default case")
  // }

  switch "Docker"  {
    // Not all the case below this will be executed like other languages. Implicit breaks in GO unlike other languages.
    case "Docker" : fmt.Println("Fallthrough Docker")
      fallthrough // All case will need fallthrough except last one
    case "Deep" : fmt.Println("Fallthrough Deep")
    case "Dive" : fmt.Println("Fallthrough Dive")
    default : fmt.Println("Default case")
  }

  switch "Docker"  {
    // Not all the case below this will be executed like other languages. Implicit breaks in GO unlike other languages.
    case "Docker", "Deep", "Dive" : fmt.Println("Fallthrough Docker1")
      fallthrough // All case will need fallthrough except last one
    case "Docker1", "Deep1", "Dive1" : fmt.Println("Fallthrough Docker2")
    case "Docker2", "Deep2", "Dive2" : fmt.Println("Fallthrough Docker3")
    default : fmt.Println("Default case")
  }

  _, err := os.Open("HelloWorldgo")
  if err != nil {
    fmt.Println(err)
  }

}
