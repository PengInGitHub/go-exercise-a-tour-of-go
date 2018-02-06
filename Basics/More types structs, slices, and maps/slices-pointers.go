package main 

import (
"fmt"

) 

func PrintSlicePointer(){
    PositionArray := [4]string{"Backend Developer", "Microservice Architecturer",
                               "Machine Learning Intern", "Frontend Developer"}
    BackendSlice := PositionArray[0:2]
    DataSlice := PositionArray[1:3]
    DataSlice[0] = "Data Mining"
    fmt.Println(DataSlice,PositionArray,BackendSlice)

}
