package main

import (
 "fmt"
 "unicode"
)

func main() {
 var letter string

 fmt.Print("Enter a letter: ")
 fmt.Scan(&letter)
 res := isLowerCase(letter)
 fmt.Println("result: ", res)
}

func isLowerCase(str string) bool {
 	for _, r := range str {
  		if !unicode.IsLower(r) {
   			return false
  		}
 	}
 	return true
}