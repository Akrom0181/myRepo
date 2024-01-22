package main

import "fmt"

func main(){
	var n int

	fmt.Println("son kiriting")
	fmt.Scan(&n)

	for i := n; i >= 1; i-- {
		if i % 2 == 1{
		fmt.Println(i)
		}
	}
}