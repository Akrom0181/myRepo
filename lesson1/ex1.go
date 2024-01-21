package main

import "fmt"

func main(){
	var(
		a int = 5
		b int = 7
		
	)
	a, b = b, a
// c = a
// a = b
// b = c
fmt.Println(a, b)

}