package main 

import (
"fmt"
"math"
) 

func PrintTypeConversion(){
    a, b := 32, 46
    var c float64 = math.Sqrt(float64(a*a+b*b))
    var d uint = uint(c)
    fmt.Println(a,b,c,d)
}
