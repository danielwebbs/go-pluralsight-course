package main

import(
  "fmt"
  "time"
  "runtime"
)

func main() {
  runtime.GOMAXPROCS(8)
  go abcGen()

  println("This comes first!")

  time.Sleep(100 * time.Millisecond)
}

func abcGen() {
  for i := byte('a'); i <= byte('z'); i++ {
  go fmt.Println(string(i))
  }
}
