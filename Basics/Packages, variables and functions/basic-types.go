package main 

import (
"fmt"

) 

var (
    TOBE bool = false
    MaxInt uint64 = 1<<64 - 1 
)

func PrintBasicType(){
    fmt.Printf("The type: %T and the value: %v\n", TOBE, TOBE)
    fmt.Printf("The type: %T and the value: %v\n", MaxInt, MaxInt)
}
