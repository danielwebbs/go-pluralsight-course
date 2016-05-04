package main

import "fmt"

func main() {
  for i:=0; i < 5; i++ {
    fmt.Println(i)
  }

  j := 0

  for {
    fmt.Println(j)

    j++

    if
    j > 5 {
      break
    }
  }

//Slice
s := []string{"foo", "bar", "buz"}

//Range says that this is an array or slice and sets the index into idx and variable into v
//This is almost like a foreach with the index
for idx, v := range s {
  fmt.Println(idx, v)
}

//Map
m := make(map[string]string)
m["first"] = "foo"
m["second"] = "bar"
m["third"] = "buz"

//Range with a Map sets k to key, and v to the value
for k, v := range m {
  fmt.Println(k, v)
}
}
