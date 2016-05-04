package main

import "fmt"

func main() {
  //Array with a set size
  myArray := [3]int{}
  myArray[0] = 42
  myArray[1] = 45
  myArray[2] = 46

  fmt.Println(myArray)

//GO will work out the size from the init values

  autoArray := [...]int{1, 2, 3}
  fmt.Println(autoArray)
  fmt.Println(len(autoArray))


  //Slice is a subset of an array which is used more than arrays in go as it manages the underline array

    mySlice := myArray[:]

    fmt.Println(myArray)
    fmt.Println(mySlice)
    mySlice = append(mySlice, 100)
    fmt.Println(mySlice)

    //Can create a slice without providing an array
    //Autosize is only good if you dont change the value often as it has to copy the array each time
    soloSlice := []float32{14., 15., 19.}
    fmt.Println(soloSlice)
    fmt.Println(len(soloSlice))

    //Make a slice with a given size
    //sizeSlice := make([]float32, 100)

    //The size of the slice is 100 and the underline array is 1000
    //arraySizeSlice := make([]float32,100,1000)

    //Maps hold mulitple values accessed by keys, can be arb data types
    //This map holds string values with ints as the keys
    myMap := make(map[int]string)

    fmt.Println(myMap)

    myMap[52] = "Foo"

    fmt.Println(myMap)
    fmt.Println(myMap[52])
    // Go returns nothing if there is no value, not null
    fmt.Println(myMap[99])
}
