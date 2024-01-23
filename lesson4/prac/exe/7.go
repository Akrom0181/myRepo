package main

import (
 "fmt"
)

func calProd(N int) int {
 	if N < 0 {
  		return -1 
 	}
	product := 1
 	for i := 1; i <= N; i++ {
  		product *= i
 	}
 	return product
}

func main() {
 	num := 5
 	result := calProd(num)
 	fmt.Println(result)
}