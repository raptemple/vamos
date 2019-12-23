package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %q \t %b \t %#X \n", i, i, i, i)
		time.Sleep(time.Second)
	}
}
