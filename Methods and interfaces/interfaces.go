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

type MyFloat float64

func (f MyFloat) Abs float64{
    if f<0 {
        float64(-f)
    }
    return float64(f)
}

type VertexInterface struct{
    X, Y float64
}

func (v *VertexInterface) Abs float64{
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func PriPrintInterfacent(){
    var a Abser

    f := MyFloat(-math.Sqrt2)
    v := VertexInterface{3,4}

    a = f //a MyFloat implements Abser
    a = &v //a My*VertexInterfaceFloat implements Abser
    
    fmt.Println()
}
