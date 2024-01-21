package main

import "fmt"

func main(){
	var(
		a float64 = 3
		b float64 = 4 
		c float64 = 5
	)
	var v = a * b * c
	var s = 2 * (a * b + b * c + a * c)
	fmt.Println(v, s)
}
