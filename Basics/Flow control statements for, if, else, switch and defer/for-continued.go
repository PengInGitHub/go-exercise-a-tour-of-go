package main 

import (
"fmt"

) 

func PrintForContinue(){
    sum:=1
    for ; sum<1000;{
        sum+=sum
    }
    fmt.Println(sum)

}
