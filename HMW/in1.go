package main

import "fmt"

func main(){

	var(
		a float64
		b float64
	)
fmt.Println("Input a side:")
fmt.Scan(&a)

fmt.Println("Input b side:")
fmt.Scan(&b)

s := a * b
fmt.Println("S =", s)
}