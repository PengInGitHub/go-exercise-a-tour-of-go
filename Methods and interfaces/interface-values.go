package main 

import (
"fmt"

)
//interface values can be thought of as a tuple of a value and a concrete type
//(value,type)

//Calling a method on an interface value executes the method of the same name on its underlying type.

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

func Print(){
    
    
    fmt.Println()
}
