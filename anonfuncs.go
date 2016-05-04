package main

import "fmt"

func main() {

addFunc := func (terms ...int)(numTerms int, sum int)  {
    for _, term := range terms {
      sum += term
    }

    numTerms = len(terms)
    //Most common use is to return the value along with a error message
    return
  }
  numTerms, sum := addFunc(1, 2, 3)
  fmt.Println("Added: ",numTerms,"terms for a total of ", sum)
}
