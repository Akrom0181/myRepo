package main

import "fmt"

func main(){
	var a int

	fmt.Println("Input number:")
	fmt.Scan(&a)

	if a % 3 == 0 && a % 5 == 0 {
		fmt.Println("FizzBuzz")
	}else if a % 3 == 0 {
		fmt.Println("Fizz")
	}else if a % 5 == 0 {
		fmt.Println("Buzz")
	}
}