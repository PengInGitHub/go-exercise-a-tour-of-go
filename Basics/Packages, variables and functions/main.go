package main

import (
	"fmt"
)

//named return values:
//a return statement without arguments return the named return values 

func split(a int) (b, c int){
   b = a * 3 / 2
   c = a + b
   return
}

func main() {
	fmt.Println(split(7))
}