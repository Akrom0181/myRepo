package main

import "fmt"

func countNum(numbers []int) (int, int, int, int, int) {
 	positives := 0
 	negatives := 0
 	zeros := 0
 	evens := 0
 	odds := 0

 	for _, num := range numbers {
 		 if num > 0 {
   			positives++
  		} else if num < 0 {
   			negatives++
  		} else {
  			zeros++
		}

  		if num%2 == 0 {
   			evens++
  		} else {
   			odds++
  		}
 	}

return positives, negatives, zeros, evens, odds
}

func main() {
 	numbers := []int{0, -1, 2, -3, 4, 0, 5, -6, 7, 0}
 	positives, negatives, zeros, evens, odds := countNum(numbers)

 	fmt.Println("Positives:", positives)
 	fmt.Println("Negatives:", negatives)
 	fmt.Println("Zeros:", zeros)
 	fmt.Println("Evens:", evens)
 	fmt.Println("Odds:", odds)
}