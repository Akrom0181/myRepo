package main

import "fmt"

func main() {
 	for i := 10; i <= 99; i++ {
  		if isP(i) {
   			fmt.Println(i)
  		}
 	}
}

func isP(n int) bool {
 	if n <= 1 {
  		return false
 	}
 	for i := 2; i*i <= n; i++ {
  		if n % i == 0 {
   			return false
  		}
 	}
 	return true
}