package main

import "fmt"

func swap(numbers []int) {
	first := numbers[0]
	last := numbers[len(numbers)-1]

	numbers[0] = last
	numbers[len(numbers)-1] = first
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	swap(numbers)

 	fmt.Println("Swapped:", numbers)
}