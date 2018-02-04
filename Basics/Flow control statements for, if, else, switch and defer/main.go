package main

import "fmt"

//to notem different from Java or C, there is no parentheses around the three components
//of the for statement, and the braces are always required

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
