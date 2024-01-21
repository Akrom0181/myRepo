package main

import "fmt"

func main(){
	var a = 764
	var yuzlik = a / 100
	var onlik = a / 10 % 10
	var birlik = a % 10

	fmt.Println(birlik * 100 + yuzlik * 10 + onlik)
}