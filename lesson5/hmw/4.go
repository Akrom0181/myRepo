
package main

import "fmt"

func SecondSmallest(numbers []int) int {
	smallest := numbers[0]
	secondSmallest := numbers[1]

 	for _, num := range numbers {
  		if num < smallest {
   			secondSmallest = smallest
   			smallest = num
  		} else if num < secondSmallest && num != smallest {
   			secondSmallest = num
 		 }
 	}

return secondSmallest
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	secondSmallest := SecondSmallest(numbers)

 	fmt.Println("Second smallest:", secondSmallest)
}