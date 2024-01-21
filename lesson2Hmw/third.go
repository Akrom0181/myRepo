package main

import (
 "fmt"
)

func main() {
 	var letter1, letter2, letter3 string

 	fmt.Print("Enter the first letter: ")
 	fmt.Scan(&letter1)

 	fmt.Print("Enter the second letter: ")
 	fmt.Scan(&letter2)

 	fmt.Print("Enter the third letter: ")
 	fmt.Scan(&letter3)

	var a = letter1 <= letter2 && letter2 <= letter3
	
	fmt.Println(a)


}