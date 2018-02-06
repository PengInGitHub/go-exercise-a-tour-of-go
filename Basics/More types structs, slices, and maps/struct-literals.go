package main 

import (
"fmt"

) 

//listing the values of struct fields
type car struct{
     brand string
     car_type string
     price float64
}
func PrintCar(){
    AudiA3 := car{"Audi", "family", 24500}
    BMWX6 := car{"BMW", "crossover",76500}
    MercedesCClass := car{brand:"Mercedes",price:37800}
    fmt.Println(AudiA3,BMWX6,MercedesCClass)

}
