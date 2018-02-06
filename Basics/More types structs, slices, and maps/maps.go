package main 

import (
"fmt"

) 
//A map maps keys to values.

type Map2D struct{
    lat, long float64
}

//make function returns a map of the given type

var m map[string]Map2D

func PrintMap(){
    m = make(map[string]Map2D)
    m["Berliner Dom"] = Map2D{52.0002, 12.32} 
    fmt.Println(m["Berliner Dom"])
}
