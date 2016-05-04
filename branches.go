package main

import "fmt"

func main() {
  //foo := 1
//This syntax does not look good, rather dont init in the if statement
  if foo := 2; foo == 1 {
    fmt.Println("bar")
  } else {
    fmt.Println("foo")
  }

switch foo := 1; foo {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
}


//Another way to do switch

foo := 5

switch {
case foo == 1:
    fmt.Println("One")
case foo > 1:
    fmt.Println("Larger")
}

}
