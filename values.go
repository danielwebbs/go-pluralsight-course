package main

import "fmt"

func main() {
  numTerms, sum := add(1 , 2, 3)
  fmt.Println("Added: ",numTerms,"terms for a total of ", sum)
}

//numTerms and sum are set to 0 and the function will know to return them
func add(terms ...int)(numTerms int, sum int)  {
  for _, term := range terms {
    sum += term
  }

  numTerms = len(terms)
  //Most common use is to return the value along with a error message
  return
}
