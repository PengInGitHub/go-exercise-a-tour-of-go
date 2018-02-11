package main 

import (
"fmt"
"math"
)

//define interface
//An interface type is defined as a set of method signatures.

//A value of interface type can hold any value that implements those methods.

type Abser interface{
    Abs() float64
}

type MyFloat2 float64

func (f MyFloat2) Abs() float64{
    if f<0 {
       return float64(-f)
    }
    return float64(f)
}

type VertexInterface struct{
    X, Y float64
}

func (v *VertexInterface) Abs() float64{
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func PrintInterface(){
    var a Abser

    f := MyFloat2(-math.Sqrt2)
    v := VertexInterface{3,4}

    a = f //a MyFloat2 implements Abser
    a = &v //a My*VertexInterfaceFloat implements Abser
    
    fmt.Println(a.Abs())
}
