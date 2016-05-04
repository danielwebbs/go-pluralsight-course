package main

import "fmt"

func main() {
message := "Hello"

  sayHello(&message)

  fmt.Println(message)


//Variadic parameters 



}
//Pointers seem to be the same as c++
func sayHello(message *string)  {
  fmt.Println(*message)

  *message = "Hello Go"
}
