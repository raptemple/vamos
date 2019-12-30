package main

import "fmt"

// switch on types
// -- normally we switch on value of variable
// -- go allows you to switch on type of variable

type contact struct {
  greeting string
  name string
}

// SwitchOnType works with interfaces
func SwitchOnType(x interface{}) {
  switch x.(type) { // this is an assert; "x is of this type"
  case int:
    fmt.Println("int")
  case string:
    fmt.Println("string")
  case contact:
    fmt.Println("contact")
  default:
    fmt.Println("unknown")
  }
}

func main() {
  SwitchOnType(7)
  SwitchOnType("Subin")
  var t = contact{
    greeting: "Merry Xmas",
    name: "Money Penny",
  }
  SwitchOnType(t)
  SwitchOnType(t.greeting)
  SwitchOnType(t.name)
}
