package main 

import (
"fmt"

)
//functions with a pointer argument must take a pointer
//var v Vertex
//ScaleFunc(&v, 10)

//in comparison methods with pointer receiver take either a pointer or a value
//var v Vertex
//&v.Scale(10) --- this is ok
//v.Scale(10) --- this is fine as well

//the reason is, Go interprets v.Scale(5) as &v.Scale(5)
//since (v *Vertext)Scale() method has a pointer receiver

//this is method pointer indirection

type VertextIndirection struct{
    a, b float64
}

func (v *VertextIndirection) ScaleMethodIn(f float64){
    v.a = v.a * f
    v.b = v.b * f
}

func ScaleFuncIn (v *VertextIndirection, f float64){
    v.a = v.a * f
    v.b = v.b * f
}

func PrintIndirection(){
    
    v := VertextIndirection{3,4}
    p := &v
    p.ScaleMethodIn(5)

    ScaleFuncIn(&v, 10)
    fmt.Println(v)
}
