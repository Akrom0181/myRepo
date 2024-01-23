package main

import (
 "fmt"
)

func printOddNumbers(N int) {
 	for i := 1; i < N; i++ {
  		if i % 2 != 0 {
   			fmt.Print(i, " ")
  		}
 	}
fmt.Println()
}

func main() {
 	printOddNumbers(5)
 	printOddNumbers(6)
}