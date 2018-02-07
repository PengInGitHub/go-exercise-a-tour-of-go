package main 

import (
"fmt"

) 
//Go supports anonymous functions, which can form closures.
//Anonymous functions are useful when you want to define a function inline without having to name it.

//define closures
//A closure is a function value that references variables from outside its body.
//The function may access and assign to the referenced variables

func add() func(int) int{
    sum := 0
    return func(x int) int{
        sum += x
        return sum
    }
}

func PrintClosures(){

    pos, neg := add(), add()
    for i:=0;i<10;i++{
        fmt.Println(
           pos(i),
           neg(-2*i),
         )
    }

}
