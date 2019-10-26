package main

import (
  "fmt"
  "time"
)

// Publish prints text to stdout after the given time has expired.
// It doesn’t block but returns right away.
func Publish(text string, delay time.Duration) {
    go func() {
        time.Sleep(delay)
        fmt.Println("BREAKING NEWS:", text)
    }() // Note the parentheses. We must call the anonymous function.
}


func main() {
  list_num := []int{1,2,3,4,5,6,7,8}
  for index, num := range list_num {
    fmt.Println(index, num)
    Publish("A goroutine starts a new thread.", 5*time.Second)
  }

  fmt.Println("Let’s hope the news will published before I leave.")
  //
  // // Wait for the news to be published.
  // time.Sleep(10 * time.Second)

  fmt.Println("Ten seconds later: I’m leaving now.")
}

// func main() {
//     go fmt.Println("Hello from another goroutine")
//     fmt.Println("Hello from main goroutine")
//
//     time.Sleep(time.Second) // give the other goroutine time to finish
// }
