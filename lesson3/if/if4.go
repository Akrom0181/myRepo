package main

import "fmt"

func main()  {
	var name string
	var lastName string
	var age int

	fmt.Println("input your name:")
	fmt.Scan(&name)

	fmt.Println("input your lastName:")
	fmt.Scan(&lastName)
	
	fmt.Println("input your age:")
	fmt.Scan(&age)

	if age % 2 == 0 {
		fmt.Println(lastName)
	}else if age % 2 == 1 {
		fmt.Println(name)
	} 
}