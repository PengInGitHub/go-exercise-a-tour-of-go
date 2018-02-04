package main

import "fmt"

//in for loop, the init and post statements are optional

func main() {
	sum := 1
	for ; sum < 1000;  {
		sum += sum
	}
	fmt.Println(sum)
}
