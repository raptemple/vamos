package main

import (
  "fmt"
)

func main() {
  fmt.Println("Data structures example")

  // 1. go Arrays have fixed size and this cannot be changed
  var array [5]int // array of integers with 5 positions ('var' keyword to initialized with zero values)
  var arrayWithValues [2]string = [2]string{"Hello,", "World"} // Initialized and values assigned

  for i := 0; i < len(array); i++ {
		array[i] = i
	}

  fmt.Println("array: ", array)
  fmt.Println("array2: ", arrayWithValues)

  // 2. Slices are associated with one underlying array and must respect it size
  var slice []int64 = make([]int64, 5)
  for i := 0; i< len(slice); i++ {
    slice[i] = int64(i)
  }
  fmt.Println("slice: ", slice)

  // Slices can grow by using the function append(), append returns a copy of the previous slice with new elements in it
	// the original slice is not modifyed, Pass by Value
	var newSlice = append(slice, 5, 6)
	fmt.Println("slice:", slice)
	fmt.Println("newSlice:", newSlice)

  slice = append(slice, -1)
	fmt.Println("slice:", slice)

  // 3. Maps need to me initialised calling make function, no var usage here
  mapA := make(map[string]int64)
  mapA["key"] = 101
  fmt.Println("mapA: ", mapA)

  mapB :=make(map[string]string)
  mapB["username"] = "raptemple"
  mapB["password"] = "encrypted"

  fmt.Println("mapB", mapB)
  fmt.Println(mapB["email"]) //prints empty string since its not found (zero Value)

  // In order to test the map return let's use error handling
  value, ok := mapB["password"]
  fmt.Println(value, ok) //ok is a boolean that answers if the key exists or not

  if value, ok := mapA["key"]; ok {
    fmt.Println("values exists:", value, ok)
  }




}
