package main 

import (
"fmt"

) 

func PrintRangeContinued(){
    pow := make([]int,10)
    for i:= range pow{
         pow[i]=1<<uint(i)
    }
    fmt.Println(pow)
}
