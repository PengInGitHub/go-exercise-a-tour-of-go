package main 

import (
"fmt"

)

//how to implement interface
//There is no explicit declaration of intent, no "implements" keyword.
//A type implements an interface by implementing its methods. 

//why implicit interfaces
//Implicit interfaces decouple the definition of an interface from its implementation
//which could then appear in any package without prearrangement.

type I interface{
     M()
}

type T struct{
    S string
}

//type T implements the interface I
func (t T) M(){
    fmt.Println(t.S)
}


func Print(){
    
    
    fmt.Println()
}
