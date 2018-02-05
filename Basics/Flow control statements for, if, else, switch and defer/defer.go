package main 

import (
"fmt"

) 

func PrintHelloWorld(){
    defer fmt.Println("世界")
    fmt.Println("Hello")

}
