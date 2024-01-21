package main 

import "fmt"

func main(){
	var a, b, c int
	fmt.Println("input numbers:")
	fmt.Scan(&a, &b, &c)
	positive := 0
	if a > 0 {
		positive++
	};if b > 0 {
		positive++
	};if c > 0 {
		positive++
	}
	fmt.Println("Number of positive nums is/are", positive)
}
