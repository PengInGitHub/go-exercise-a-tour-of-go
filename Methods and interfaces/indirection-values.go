package main 

import (
"fmt"
"math"
)

func (v *VertextIndirection) AbsIn() float64{ // donot forget AbsIn()
    return math.Sqrt(v.a*v.a + v.b*v.b)
}

func AbsInFunc(v *VertextIndirection) float64{
    return math.Sqrt(v.a*v.a + v.b*v.b)
}

func PrintIndirectionValues(){
    v := VertextIndirection{3,4}
    
    fmt.Println(v.AbsIn())
}
