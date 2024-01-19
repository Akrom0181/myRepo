package main

import "fmt"

func main(){
	var(
		a = 3
		b = 4
		c = 5
		// d int
	)
	a, b, c = b, c, a
	// d = a
	// a = b
	// b = c
	// c = d
	
	fmt.Println("a", a,"b", b,"c", c)
}