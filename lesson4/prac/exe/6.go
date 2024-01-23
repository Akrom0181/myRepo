package main

import (
 "fmt"
)

func isPrime(n int) bool {
 	if n <= 1 {
  		return false
 	}
 	for i := 2; i*i <= n; i++ {
  		if n%i == 0 {
   			return false
  		}
 	}
 	return true
}

func primeDiv(num int) {
 	for i := 2; i <= num; i++ {
  		if num % i == 0 && isPrime(i) {
   			fmt.Print(i, " ")
  		}
 	}
}

func main() {
 	number := 24
 	fmt.Print("Prime divisors ", number, ": ")
 	primeDiv(number)
}