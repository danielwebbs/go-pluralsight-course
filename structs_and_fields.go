package main

import "fmt"

func main() {
/*  foo := new(myStruct)
  foo.myField = "bar" */

  //Composite literal syntax
  foo := &myStruct{"bar"}

  fmt.Println(foo.myField)
}

type myStruct struct {
  myField string
}
