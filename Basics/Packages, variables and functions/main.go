package main

import (
	"fmt"
)

//short variable declaration:
//k := 3 could be used in a function when declaring a variable
//outside a function, every statement starts with a keyword

var a, b = true, false

func main() {
    c := "Ok, fine"
    d := 32
	fmt.Println(a,b,c,d)
}