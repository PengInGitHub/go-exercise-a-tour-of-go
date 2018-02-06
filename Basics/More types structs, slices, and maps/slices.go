package main 

import (
"fmt"

) 

func PrintSlices(){
    primes := [6]int{2,3,5,7,11,13}
    var PrimesSlice []int = primes[1:4]
    fmt.Println(PrimesSlice)

}
