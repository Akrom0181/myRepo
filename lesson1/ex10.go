package main

import "fmt"

func main(){
	var a = 125
	var yuzlik = a / 100
	var onlik = a / 10 % 10
	var birlik = a % 10

	fmt.Println(onlik * 100 + birlik * 10 + yuzlik)
}