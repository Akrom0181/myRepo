package main

import "fmt"

func main(){
	res := numSum(14)
	fmt.Printf()
}

func numSum(a int) int{
	sum := 0
	for i := 1; i <= a; i++ {
		sum += i
	}
return sum
}