package main

/*This is a way of implementing a Enum*/
const (
  first = iota
  second
  third
)

/*bitshift on iota but why?*/
const (
  test = 1 << (10 * iota) // This code will execute every time for each new iota
  rock
)

func main() {
  println(first)
  println(second)
  println(third)
}
