package main

import (
  "fmt"
  "rsc.io/quote"
)

func main() {
  fmt.Println(Hello())
  fmt.Println(OptQuote()) //return quote.Opt()
  fmt.Println(GlassQuote()) //return quote.Glass()
  fmt.Println(GoQuote()) //return quote.Go()
}

func Hello() string {
    return quote.Hello()
}

func OptQuote() string {
	return quote.Opt()
}

func GlassQuote() string {
	return quote.Glass()
}

func GoQuote() string {
	return quote.Go()
}
