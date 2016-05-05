package main

import(
  "runtime"
)

func main() {
  runtime.GOMAXPROCS(8)
  doneCh := make(chan bool)
  ch := make(chan string)

  go abcGen(ch)
  go printer(ch, doneCh)

  println("This comes first!")

  <-doneCh
}

func abcGen(ch chan string) {
  for i := byte('a'); i <= byte('z'); i++ {
  ch <- string(i)
  }
  close(ch)
}

func printer(ch chan string, doneCh chan bool) {
  for i:= range ch{
    println(i)
  }

  doneCh <- true
}
