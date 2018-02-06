package main 

import (
"fmt"

) 
//slices can be created by the built-in function make
//this is how dynamically-sized arrays created

//how make works: make allocates a zeroed array and returns a slice refers to the array
//make has three parameters: array type, array length and array capacity

func PrintMake(){
    a := make([]int,5)
    PrintSliceHelper("a",a)

    b := make([]int,5,50)
    PrintSliceHelper("b",b)

    c := make([]int,0,50)
    PrintSliceHelper("c",c)

    d := b[:2]
    PrintSliceHelper("d",d)
}

func PrintSliceHelper(s string, a[]int){
    fmt.Printf("slice %s, len=%d, cap=%d, %v\n",s,len(a),cap(a),a)
}

