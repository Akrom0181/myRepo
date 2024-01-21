package main

import "fmt"

func main(){
	var a int
	fmt.Println("narxni kiriting:")
	fmt.Scan(&a)

	for i := 1; i <= 10; i++ {
		fmt.Println(i, i * a)
	}
}