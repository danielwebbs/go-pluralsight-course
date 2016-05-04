package main

import "fmt"

func main() {
  var myInt int
  myInt = 42

  fmt.Println(myInt)
  /* The period after the number is used to tell GO that the number is a float, go does not like to inplicitly convert types
  Float has no standard so you must specify the size*/
  var myFloat32 float32 = 60.

  fmt.Println(myFloat32)

/** := tells GO to get the type from he value being assigned  */
  myString := "Hello Go"
  fmt.Println(myString)

  myComplex := complex(2, 3)
  fmt.Println(myComplex)
  fmt.Println(real(myComplex))
  fmt.Println(imag(myComplex))
}
