package main

import "fmt"

func main(){
	var a float64
	fmt.Println("narxni kiriting:")
	fmt.Scan(&a)

	for i := 0.0; i <= 1.0; i += 0.1 {
		fmt.Println(i, i * a)
	}
}