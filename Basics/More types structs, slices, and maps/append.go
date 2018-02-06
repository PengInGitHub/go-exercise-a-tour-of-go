package main 

import (
"fmt"

) 

func PrintAppend(){
    var s []int
    HelperMethodPrintSlice(s)

}

func HelperMethodPrintSlice(s []int){
     fmt.Printf("len=%d, cap=%d, %v\n",len(s),cap(s),s)
}