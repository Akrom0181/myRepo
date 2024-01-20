package main

import(
 "fmt"
 "math"
)
func main(){
	var number float64

	fmt.Println("Input number:")
	fmt.Scan(&number)

	a := math.Pow(number, 3)
	fmt.Println("3rd degree of", number, "=", a)

}