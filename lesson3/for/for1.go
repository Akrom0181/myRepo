package main

import "fmt"

func main(){
	var k, n int

	fmt.Println("Son kiriting:")
	fmt.Scan(&k, &n)

	for i := 1; i <= n; i++ {
		fmt.Print(k, " \n")
	}
}