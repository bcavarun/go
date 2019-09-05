package main



import (
  "fmt"
  // "runtime"
  "os"
)

var (
  globalName = "Varun BHANDARI"
  globalCast string
  globalNum int
)

func main()  {
  // globalCast = "Hellooooooo"
  // name := "Varun Bhandari"
  // fmt.Println("Running go in ", runtime.GOOS)
  fmt.Println(globalNum)
  fmt.Println(globalCast)
  os.G
  // fmt.Println("Hello ", globalName)
  // fmt.Println("Hello ", globalCast)
  // fmt.Println(os.Environ())
  fmt.Println(os.Getenv("Name"))
  // for _, env := range os.Environ() {
  //   fmt.Println(env)
  // }
  // demoPointers(&name)
  // fmt.Println("Hello ", name)
}



func demoPointers(name *string)  {
  *name = "Srt"
  fmt.Println(*name)

}
