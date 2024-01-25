package main

import "fmt"

func SecondLar(numbers []int) int {
 	largest := numbers[0]
 	secondLar := numbers[0]

 	for _, num := range numbers {
  		if num > largest {
   			secondLar = largest
   			largest = num
  		} else if num > secondLar && num != largest {
  			secondLar = num
  		}
 	}

return secondLar
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	secondLar := SecondLar(numbers)

	fmt.Println("Second largest:", secondLar)
}