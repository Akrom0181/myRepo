package main

import "fmt"

func main(){
	var(
		a = 5
		b = 7
	)
a, b = b, a

fmt.Println("a=", a, "b=", b)
}