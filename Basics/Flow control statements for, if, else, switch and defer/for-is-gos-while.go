package main 

import (
"fmt"

) 

func PrintForAsWhile(){
    sum:=1
    for sum<1000{
         sum+=sum
    }
    fmt.Println(sum)

}
