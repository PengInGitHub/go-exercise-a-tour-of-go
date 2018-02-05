package main

import (
	"fmt"
	"math/rand"
)

func PrintRandomInt(x int) {
    fmt.Println("My favorite number is", rand.Intn(x))
}