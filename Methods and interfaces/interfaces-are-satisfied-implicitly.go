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

func Print(){
    
    
    fmt.Println()
}
