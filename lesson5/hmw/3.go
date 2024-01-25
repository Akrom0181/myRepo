package main

import "fmt"

func cal(numbers []int) (int, int) {
 	evenIn := 1
 	oddIn := 0

	for i, num := range numbers {
  		if i%2 == 0 {
   			evenIn *= num
  		} else {
   		 	oddIn += num
  		}
 	} 

return evenIn, 	oddIn
}

func main() {
 	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
 	evenIn, oddIn := cal(numbers)

	fmt.Println("evenIn:", evenIn)
	fmt.Println("oddIn:", oddIn)
}