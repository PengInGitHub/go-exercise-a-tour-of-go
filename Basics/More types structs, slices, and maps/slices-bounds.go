package main 

import (
"fmt"

) 

func PrintSliceExample(){
    SliceExample := []int{19,12,31,5,21,4}
    SliceExampleSlice1 := SliceExample[:3]
    SliceExampleSlice2 := SliceExample[3:]
    fmt.Println(SliceExample,SliceExampleSlice1,SliceExampleSlice2)

}
