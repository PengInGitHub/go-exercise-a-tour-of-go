package main 

import (
"fmt"

)

//declare methods with pointer receivers
//receiver has literal types *T
//methods with pointer receiver could modify the value the receiver points 

//pointer receivers are more common than value receivers, because methods often modify their receivers

func (r RightTriangle) Scale(f float64){
    r.SideA = r.SideA * f 
    r.SideB = r.SideB * f 
}

func PrintMethodsPointers(){
    
    r := RightTriangle{300,400}
    r.Scale(10)
    fmt.Println(SideCLength(r))
}
