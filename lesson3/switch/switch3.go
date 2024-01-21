package main

import "fmt"

func main(){
	fmt.Println("input num:")
	var month int
	fmt.Scan(&month)
 	switch month{
 		case 1, 2, 12:
  			fmt.Println("Winter")
 		case 3, 4, 5: 
  			fmt.Println("Spring")
 		case 6, 7, 8: 
  			fmt.Println("Summer")
 		case 9, 10: 
  			fmt.Println("Autumn")  
 		default:
  			fmt.Println("You can input between 1 and 12!")  
 	}
}