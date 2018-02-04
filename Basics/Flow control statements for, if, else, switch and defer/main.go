package main

import "fmt"
import "math"

//if
//if in Go is same like for, no () is required for statements but {} are required

func sqrt(a float64) float64 {

   if a<0 {
        return sqrt(-a)
    }
    return math.Sqrt(a)
}

func main() {
	fmt.Println(sqrt(-2), sqrt(16))
}
