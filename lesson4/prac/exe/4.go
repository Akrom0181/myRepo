package main

import "fmt"

func main() {
	result := addSum(5)
	fmt.Println(result)
	
}

func addSum(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}