package main

import (
	"fmt"
)

//a function can return any number of results

func swap(a , b  string) (string, string){
   return b, a
}

func main() {
	fmt.Println(swap("Hello", "World"))
}