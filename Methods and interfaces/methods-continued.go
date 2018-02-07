package main 

import (
"fmt"

)

//declare a method on non-struct type

type MyFloat float64

func (m MyFloat) Abs() float64{
    if m< 0 {
        return float64(-m)
    }
    return float64(m)
}

func PrintMyFloat(){
    var m MyFloat = -64   
    fmt.Println(m.Abs())
}
