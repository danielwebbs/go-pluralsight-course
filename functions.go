package main

import "fmt"

func main() {
message := "Hello"

  sayHello(&message)

  fmt.Println(message)

 sayHelloV("Tod", "Rob")


}
//Pointers seem to be the same as c++
func sayHello(message *string)  {
  fmt.Println(*message)

  *message = "Hello Go"
}

//Variadic parameters
func sayHelloV(messages ...string)  {
  for _, message := range messages {
    fmt.Println(message)
  }
}
