package main

import (
  "fmt"
)

func main() {
  fmt.Println(multiple_returns())
  fmt.Println(multiple_args(1, 2, 3))

  clojure := func() string {
    return "10110101010101"
  }
  fmt.Println(clojure())

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
