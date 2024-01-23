package main

import "fmt"

func main() {
third(1, 0, 4)
	
}

func third(a, b, c int) {
	if a > b && b > c {
		fmt.Println(b)
	}else if b > c && c > a {
		fmt.Println(c)
	}else if c > a && a > b {
		fmt.Println(a)
	}
}