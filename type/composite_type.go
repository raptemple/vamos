package main

import (
  "fmt"
)

type (
  MyInt int
  Age int
  Text string
)

type IntPtr *int

type Book struct {
  author string
  title string
  pages int
}

type Convert func(i0 int, i1 bool) (out0 int, out1 string)
type StringArray [5]string
type StringSlice []string

func f() {
  // define in func, it can only be used in func
  type PersonAge map[string]int
  type MessageQueue chan string
  type Reader interface{
    Read([]byte) int
  }
}


func main() {
  a := StringSlice{"1", "2", "3"}
  fmt.Println(a[:])
}
