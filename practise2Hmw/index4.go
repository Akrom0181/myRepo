package main

import(
	"fmt"
	"math"
)

func main(){
	var diametr float64
	
	fmt.Println("Diametrni kiriting:")
	fmt.Scan(&diametr)

	l := math.Pi * diametr
	fmt.Println(l)
}