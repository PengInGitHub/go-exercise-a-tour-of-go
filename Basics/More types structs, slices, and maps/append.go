package main 

import (
"fmt"

) 

func PrintAppend(){
    var s []int //nil slices
    HelperMethodPrintSlice(s)

    //append works on nil, len=1, cap=1
    s=append(s,0)
    HelperMethodPrintSlice(s)

    //slice keeps growing
    s=append(s,10)
    HelperMethodPrintSlice(s)

    //slice keeps growing
    s=append(s,10)
    HelperMethodPrintSlice(s)

}

func HelperMethodPrintSlice(s []int){
     fmt.Printf("len=%d, cap=%d, %v\n",len(s),cap(s),s)
}