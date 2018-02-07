package main 

import (
"fmt"

) 


func PrintMapMoreConcise(){
    var m = map[string]MapLatLong{
        "Friedrich":{40.1,21.34},
        "Friedrich2":{10.1,41.35534},
    }
    fmt.Println(m)

}
