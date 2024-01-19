package main

import( 
	"fmt"
    "math"
)

func  main(){
	var a float64

	fmt.Println("Kvadratni tomonini kiriting: ")
	fmt.Scan(&a)

	s := math.Pow(a, 2)
	fmt.Println(s)
}