package main 

import (
"fmt"

) 

type library struct{
    location string
    closed_time int
}

func PrintLibrary(){
    hu_bib := library{"Friedrichstr", 24}
    fmt.Println(hu_bib)
    p := &hu_bib
    p.closed_time = 22
    fmt.Println(hu_bib)
}
