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

func PriPrintInterfacent(){
    
    
    fmt.Println()
}
