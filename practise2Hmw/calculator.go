package main

import "fmt"

func main(){
	var(
		a int
		b int
		s string
		plus = "+"
		minus = "-"
		m = "*"
		d = "/"
	)
	fmt.Println("input first num:")
	fmt.Scan(&a)

	fmt.Println("input second num:")
	fmt.Scan(&b)
	
	fmt.Println("input command")
	fmt.Scan(&s)

	if s == plus{
		fmt.Println(a + b)
	}else if s == minus{
		fmt.Println(a - b)
	}else if s == m {
		fmt.Println(a * b)
	}else if s == d && b != 0{
		fmt.Println(a / b)
	}
}