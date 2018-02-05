package main 

import (
"fmt"

) 

var parrot_cute int = 32 

func PrintVarWithInit(){
   var cat, dog, lizard = false, true, "Ok fine"
   fmt.Println(parrot_cute, cat, dog, lizard)
}
