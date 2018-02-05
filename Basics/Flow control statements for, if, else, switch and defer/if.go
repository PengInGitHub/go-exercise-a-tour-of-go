package main 

import (
"fmt"
"math"
) 

func sqrt(n float64) float64{
    if n<0 {
        return sqrt(-n)
    }
    return math.Sqrt(n)
}

func PrintSqrt(n float64){
    
    fmt.Println(sqrt(n))

}
