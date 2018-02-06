package main 

import (
"fmt"

) 

type phone struct{
    brand string
    manufactor string
}

func PrintStruct(){
    
    fmt.Println(phone{"P10", "Huawei"})
    fmt.Println(phone{"iPhoneX", "Apples"})
    M6 := phone{brand:"M6",manufactor:"XiaoMi"}
    fmt.Println(M6.brand)
    p := &M6
    fmt.Println(p.manufactor)
    p.brand = "XiaoMi Mix" 
    fmt.Println(p.brand)
}
