package main 

import (
    "fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//method in Go is the same as method in OOP, which belongs to a class
//and could be called only from an instance of the class

func PrintMethod(){
    v := Vertex{3, 4}
	fmt.Println(v.Abs())
    
}
