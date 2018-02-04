package main

import (
	"fmt"
)

//variables with initializers:
//type of variable could be omitted if initializer is present

var a, b, c = true, false, "Oh Yes!"

func main() {
    var d = 32
	fmt.Println(a,b,c,d)
}