package main

import "fmt"

func main(){
	var a, b int
	fmt.Println("son kiriting:")
	fmt.Scan(&a, &b)
	count := 0
	if a > b {
	for i := a-1; i > b; i-- {
		count++
		fmt.Println(i)
	}
	fmt.Println("soni",count)
	}else {
		for i := b-1; i > a; i-- {
			count++
			fmt.Println(i)
		}
		fmt.Println("soni",count)
}
}