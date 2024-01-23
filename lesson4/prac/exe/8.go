package main

import "fmt"

func check(a, b, c int) int {
 	if a+b == c {
  		return 1
 	}
 		return 0
}

func main() {
 	result1 := check(2, 3, 4)
 	fmt.Println("Result for (2, 3, 4):", result1)

 	result2 := check(1, 2, 3)
 	fmt.Println("Result for (1, 2, 3):", result2)
}