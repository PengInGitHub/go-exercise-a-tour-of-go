package main 

import (
"fmt"

)

//reasons to choose pointer receiver
//1.methods can modify the value the receiver points to
//2.avoid copying value in each method call, especially efficient when the receiver is a large struct


func Print(){

    fmt.Printf("a")
}
