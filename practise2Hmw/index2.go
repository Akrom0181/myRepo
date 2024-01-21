package main

import "fmt"

func main(){
	var(
		a float64
		b float64
		c float64
	)
	fmt.Println("A tomonini kiriting:")
	fmt.Scan(&a)

	fmt.Println("B tomonini kiriting:")
	fmt.Scan(&b)

	fmt.Println("C tomonini kiriting:")
	fmt.Scan(&c)
	
	Condition := a + b > c && a + c > b && b + c > a
	fmt.Println(Condition)

}