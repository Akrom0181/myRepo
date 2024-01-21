package main

import "fmt"

func main(){
	var a int
	fmt.Println("input year:")
	fmt.Scan(&a)
	if (a % 4 == 0 && a % 100 != 0) || a % 400 == 0 {
		fmt.Println("Leapyear")
	}else {
		fmt.Println("year is not leapyear")
	}
}
