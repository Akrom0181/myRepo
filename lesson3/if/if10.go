package main

import "fmt"

func main(){
	var a, b int
	fmt.Println("Input 2 numbers:")
	fmt.Scan(&a, &b)
	if a > b {
		fmt.Println(a)
	}else if a < b {
		fmt.Println(b)
	}
}