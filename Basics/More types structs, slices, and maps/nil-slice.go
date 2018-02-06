package main 

import (
"fmt"

) 

func PrintNilSlice(){
    var NilSlice []int //why "NilSlice := []int" does not work?
    fmt.Println(NilSlice, len(NilSlice), cap(NilSlice))
    if NilSlice == nil {
        fmt.Printf("It is nil.")
    }    

}
