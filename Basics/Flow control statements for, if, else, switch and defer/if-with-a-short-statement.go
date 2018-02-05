package main 

import (
"math"
) 

func pow(n,power,limit float64) float64{
    if v:=math.Pow(n,power);v<limit{
        return v
    }
    return limit
}
