package main 

import (
"fmt"

) 

//why map?
//One of the most useful data structures in computer science is the hash table. Many hash table implementations exist with varying 
//properties, but in general they offer fast lookups, adds, and deletes.

//map in golang
//Go provides a built-in map type that implements a hash table.

//what does map look like in golang
//map[KeyType]ValueType

//example
//var m map[string]int
//map is reference type, same as slice or pointer, so m is nil and doesnot point to an initialized map

//why use "make" function
//write directly to an nil map will cause a runtime panic, so "var m map[string]int" is not a proper way to init a map
//to do that, use "make": m=make(map[string]int)

//how "make" works
//"make" allocates and initializes a hash map data structure,
//and return a map value points to the hash maps data structure

func PrintMakeMap(){
    m := make(map[string]int)

//working with maps
    m["Answer"]=42
    fmt.Println("The value:",m["Answer"])

}
