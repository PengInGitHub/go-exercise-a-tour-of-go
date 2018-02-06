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

    //players take turns
    board[0][0]="x"   
    board[2][2]="o"    
    board[0][2]="x"    
    board[1][0]="o"    
    board[1][2]="x"    
 
    fmt.Println(board)

}
