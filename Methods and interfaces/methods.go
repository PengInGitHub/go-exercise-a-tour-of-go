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

//add another example for method

type TheInteger int

func (a TheInteger) AddInt(b int) int {
     return int(a) + b
}

func PrintMethod(){
    var x TheInteger = 100 
	fmt.Println(x.AddInt(26))
    
}
