package main

import "fmt"

func main(){
	fmt.Println("input num:")
	var month int
	fmt.Scan(&month)
 	switch month{
 		case 1, 3, 5, 7, 8, 9, 11, 12:
  			fmt.Println("31 kundan iborat")
 		case 2: 
  			fmt.Println("28 yoki 29 kundan iborat")
 		case 4, 6, 10: 
  			fmt.Println("30 kundan iborat") 		
  		default:
			fmt.Println("You can input between 1 and 12!")  
 	}
}