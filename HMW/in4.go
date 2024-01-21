package main

import "fmt"

func main(){
	var a, b float64

	fmt.Println("a va b ni kiriting:")
	fmt.Scan(&a, &b)
	
	p := (a + b) * 2
	fmt.Println(p)
}