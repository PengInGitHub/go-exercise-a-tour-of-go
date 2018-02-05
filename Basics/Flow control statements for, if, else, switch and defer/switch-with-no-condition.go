package main 

import (
"fmt"
"time"
) 

func PrintGreeting(){
    
    today := time.Now()
    switch {

    case today.Hour() < 12:
        fmt.Println("Morning!")

    case today.Hour() < 17:
        fmt.Println("Good Afternoon!")

    default:
        fmt.Println("Good Evening!")

    }

}
