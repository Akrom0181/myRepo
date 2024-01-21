package main 

import "fmt"

func main(){
	var a, b int

	fmt.Println("input 2 num:")
	fmt.Scan(&a, &b)

	if a > b {
		fmt.Println(a, b)
	}else if a < b {
		fmt.Println(b, a)
	}
}