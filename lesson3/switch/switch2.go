package main

import "fmt"

func main(){
	fmt.Println("input num:")
	var grade int
	fmt.Scan(&grade)
 	switch grade{
 		case 1:
  			fmt.Println("yomon")
 		case 2: 
  			fmt.Println("qoniqarsiz")
 		case 3: 
  			fmt.Println("qoniqarli")
 		case 4: 
  			fmt.Println("yaxshi") 
		case 5: 
  			fmt.Println("a'lo")  
 		default:
  			fmt.Println("You can input between 1 and 5!")  
 	}
}