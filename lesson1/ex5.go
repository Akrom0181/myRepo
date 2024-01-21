package main

import "fmt"

func main(){
	
	number := 85
	firstN := number / 10
	sedcondN := number % 10
	
	fmt.Println(sedcondN * 10 + firstN)
}