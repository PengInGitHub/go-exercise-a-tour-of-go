package main 

import (
"fmt"
) 

const World = "世界"

func PrintConst(){
    const Pi = 3.14159
    fmt.Println("Pi: ", Pi)
    fmt.Println("Hello", World)
}
