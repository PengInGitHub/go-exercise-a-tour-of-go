package main 

import (
"fmt"
"math"
) 

func PowIfElse(n,power,limit float64) float64 {
    if v:= math.Pow(n,power); v<limit {
         return v
    }else{fmt.Printf("%g >= %g",v,limit) }
    
    return limit
}

