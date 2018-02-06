package main 

import (
"fmt"

) 

//create a tic-tac-toe board

func PrintBoard(){
    board := [][]string{
             []string{"_","_","_"},
             []string{"_","_","_"},
             []string{"_","_","_"},
    }
  
    fmt.Println(board)

}
