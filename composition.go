package main

import "fmt"

func main() {
  emp := enhancedMessagePrinter{messagePrinter{"foo"}}
  emp.message = "foo"
  emp.printMessage()
}

type messagePrinter struct {
  message string
}

func (mp *messagePrinter) printMessage() {
  fmt.Println(mp.message)
}

type enhancedMessagePrinter struct {
  messagePrinter
}
