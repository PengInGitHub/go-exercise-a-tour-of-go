package main 

import (
"fmt"
"time"
) 

func PrintSwitchOrder(){
    fmt.Printf("When is Saturday?\n")
    today := time.Now().Weekday()//Monday, Tuesday ...
    
    switch time.Saturday{
        case today + 0:
            fmt.Println("today")
             
        case today + 1:
            fmt.Println("tomorrow")  

        case today + 2:
            fmt.Println("in two days")

        default :
            fmt.Println("far away")

}
       

}
