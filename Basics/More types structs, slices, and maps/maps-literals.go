package main 

import (
"fmt"

) 
//Map literals are like struct literals, but the keys are required.

type MapLatLong struct{
    lat, long float64
}

func PrintMapLiterals(){
    var m = map[string]MapLatLong{

        "Friedrichstr":MapLatLong{
            53.2,11.2, //why one more comma after two numbers?
        },
        "Gesundbounen":MapLatLong{
            54.22,10.2,//why one more comma after two numbers?
        },
    }
    
    fmt.Println(m)

}
