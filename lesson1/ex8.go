package main

import "fmt"

func main(){
	var a = 857
	var yuzlik = a / 100
	var onlik = a / 10 % 10
	var birlik = a % 10
	var sumDigit = yuzlik + onlik + birlik

	fmt.Println(sumDigit)
}
