package main

import "fmt"

func main(){
	var(
		a float64
		b float64
	)
fmt.Println("a tomonini kiriting:")	
fmt.Scan(&a)

fmt.Println("b tomonini kiriting:")	
fmt.Scan(&b)

p := 2 * (a + b)
fmt.Println("result:",p)
}