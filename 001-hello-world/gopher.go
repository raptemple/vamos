package main

import (
  "fmt"
)

func main() {
  fmt.Println(multiple_returns())
  fmt.Println(multiple_args(1, 2, 3))

  closure := func() string {
    return "10110101010101"
  }
  fmt.Println(closure())

}


func multiple_returns() (string, string) {
  return "FirstName,", "LastName"
}

func multiple_args(args ...int) (int) {
  sum := 0
  for _, v := range args {
    sum += v
  }
  return sum
}
