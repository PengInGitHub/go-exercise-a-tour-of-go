package main 

import (
"fmt"
"math"
)

type RightTriangle struct{
    SideA, SideB float64
}

func SideCLength(r RightTriangle) float64{
    return math.Sqrt(r.SideA*r.SideA + r.SideB*r.SideB)
}

func PrintMethodFunc(){
    r := RightTriangle{30,40}
    //a method is just a function with a receiver argument
    fmt.Println(SideCLength(r))
}
