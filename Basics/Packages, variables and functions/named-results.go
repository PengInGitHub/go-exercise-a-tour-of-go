package main 

import (
//"fmt"
"math"
) 

func split(a float64) (x, y float64){
    x = math.Sqrt(a)
    y = x + a*23
    return x, y
}
