package main

import "fmt"
/* Var is assigned before the init function is run, main is then run*/
var (
  message = "Hello World"
)

func main() {
  fmt.Println(message)
}

func init()  {
  message = "Hello Go"
}
