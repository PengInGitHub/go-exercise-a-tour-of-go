package main

import (
	"fmt"
)

//zero values:
//variables are declared without explict initial values are given their zero values
//outside a function, every statement starts with a keyword

func main() {
    var a bool
    var b string
    var c int
    var d float64
	fmt.Printf("%v %q %v %v",a,b,c,d)
}