package main 

import (
"fmt"

)

type AStruct struct{
    a, b float64
}

func ScaleFunc(s AStruct, f float64){
    s.a = s.a * f
    s.b = s.b * f
}

func PrintMethodsPointersEx(){
    
    
    fmt.Println()
}
