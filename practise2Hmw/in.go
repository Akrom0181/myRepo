package main

import "fmt"

func main(){
	fmt.Println((!(true && false) || !(true || false)) && ((true && false) && !(true || false)))
	fmt.Println((111 % 1 == 1 && 1 * 0 == 0) || !(0/1 == 1) || (3695 / 2 == 1847))
}