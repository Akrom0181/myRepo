package main

import(
	"fmt"
	"math"
)
func main(){
	var a float64
	fmt.Println("radiusni kiriting:")
	fmt.Scan(&a)

	s := a * a * math.Pi 
	fmt.Print(s)
}