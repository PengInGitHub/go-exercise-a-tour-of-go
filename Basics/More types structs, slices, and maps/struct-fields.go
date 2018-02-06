package main 

import (
"fmt"

) 

type apartment struct{
    location string
    size float64
    rent float64
}

func PrintStructFields(){
    
    fmt.Println(apartment{"Tegel", 64, 600})
    licht := apartment{"Lichtenberg",26,350}
    fmt.Println(licht.size)

}
