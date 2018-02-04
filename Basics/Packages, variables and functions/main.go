package main

import (
	"fmt"
    "math"
)

//type conversion:
//the expression T(v) converts variable v to type T
//i := 64
//f := float64(i)
//u := uint(f)

func main() {
    a, b := 3, 4 
    c := math.Sqrt(float64(a*a+b*b))
    d := uint(c)
	fmt.Println(a,c,d)
}