package main 

import (
"fmt"

) 

func PrintSliceLen(){
    SliceExample := []int{19,12,31,5,21,4}
    PrintSliceFunc(SliceExample)

}


func PrintSliceFunc(s []int){
    fmt.Printf("length: %d, capacity: %d and the slice: %v\n",len(s),cap(s),s)
}