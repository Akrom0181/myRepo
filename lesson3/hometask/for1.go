package main

import "fmt"

func main(){
	var n int

	fmt.Println("son kiriting")
	fmt.Scan(&n)

	for i := n; i >= 1; i-- {
		fmt.Println(i)
	}
}