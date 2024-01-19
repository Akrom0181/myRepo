package main

import "fmt"

func main(){
	var a string
	fmt.Println("harf kiriting: ")
	fmt.Scan(&a)

	isVowel := a == "A" || a == "E" || a == "I" || a == "O" || a == "U" || a == "a" || a == "e" || a == "i" || a == "o" || a == "u" 
	fmt.Printf("Unli%v\n ", isVowel)

}
