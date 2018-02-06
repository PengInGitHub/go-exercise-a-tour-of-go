package main 

import (
"fmt"

) 

func PrintSliceLiterals(){
    IntSliceLiteral := []int{1,2,3}
    fmt.Println(IntSliceLiteral)

    BoolSliceLiteral := []bool{false,false,true}
    fmt.Println(IntSliceLiteral)

    StructSliceLiteral := []struct{
                       name string
                       height float64
                      }{
                       {"Yao",2.26},      
                       {"T-mac",1.96},      
                       {"Kobe",1.90},      
 }
    fmt.Println(IntSliceLiteral,BoolSliceLiteral,StructSliceLiteral)

}
