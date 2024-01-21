package main

import "fmt"

func main(){
	var num = 345
	var birlik = num % 10
	var onlik = num / 10 % 10

	fmt.Println(birlik, onlik)
}
