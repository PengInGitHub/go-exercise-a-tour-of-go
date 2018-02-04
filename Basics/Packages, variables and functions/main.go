package main

import (
	"fmt"
    //"math"
)

//constants:
//constants are declared with the keyword const
//constants cannot be declared with the syntax :=

const Pi = 3.14

func main() {
    const World = "World"
    const Truth = true
	fmt.Println(Pi)
    fmt.Println(World, Truth)
}