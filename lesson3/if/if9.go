package main 

import "fmt"

func main(){
	var a, b, c int
	fmt.Println("input numbers:")
	fmt.Scan(&a, &b, &c)
	positive := 0
	negative := 0
	if a > 0 {
		positive++
	};if b > 0 {
		positive++
	};if c > 0 {
		positive++
	};if a < 0{
		negative++
	};if b < 0{
		negative++
	};if c < 0{
		negative++
	}
	fmt.Println("Number of positive nums is/are", positive)
	fmt.Println("Number of negative nums is/are", negative)

}