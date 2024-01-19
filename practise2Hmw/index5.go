package main

import "fmt"

func main(){
	var num float64
	
	fmt.Println("Son kiriting:")
	fmt.Scan(&num)

	square := num * num
	fmt.Println(square)
}