package main

import (
	"fmt"
)

func add(a int, b int) int{
   return a+b
}

func main() {
	fmt.Printf("Value of 3+4: %v", add(3,4))
}