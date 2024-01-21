package main

import (
	"fmt"
	"math"
)

func main(){
	var a float64
	fmt.Println("input number:")
	fmt.Scan(&a)

	if int(a) % 2 == 0 {
		fmt.Println(math.Pow(a, 2))
	}else if int(a) % 2 == 1 {
		fmt.Println(math.Pow(a, 3))
	}
}