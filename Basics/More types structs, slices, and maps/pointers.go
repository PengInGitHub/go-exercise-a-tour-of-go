package main 

import (
"fmt"

) 

func PrintPointer(){
    a, b := 32,67
    p := &a
    fmt.Println(*p)
    *p = 311
    fmt.Println(a)

    p = &b
    *p = *p/3
    fmt.Println(b)
}
