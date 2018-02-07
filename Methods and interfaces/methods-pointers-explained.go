package main 

import (
"fmt"
"math"
)

type AStruct struct{
    a, b float64
}

func ScaleFunc(s *AStruct, f float64){
    s.a = s.a * f
    s.b = s.b * f
}

func AbsFunc(s AStruct) float64{
    return math.Sqrt(s.a*s.a + s.b*s.b)
}

func PrintMethodsPointersEx(){
    s := AStruct{3,4} 
    ScaleFunc(&s,10)//use &s
    fmt.Println(AbsFunc(s))
}
