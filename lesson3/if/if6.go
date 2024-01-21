package main

import "fmt"

func main(){

	var a int
	fmt.Println("input number:")
	fmt.Scan(&a)
	if a > 0{
		a++
		fmt.Println(a)
	}else {
		a-=2
		fmt.Println(a)
	}
}