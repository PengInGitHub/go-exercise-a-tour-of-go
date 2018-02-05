package main 

import (
"fmt"

) 

func PrintSum() {
    sum := 0
    for i:=0; i<100; i++ {
        sum+=i
    }
    fmt.Println(sum)
}
