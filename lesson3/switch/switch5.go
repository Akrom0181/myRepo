package main

import "fmt"

func main(){
	fmt.Println("input num:")
	var a, b int 
	var amal int
	fmt.Println("a va b ni kiriting:")
	fmt.Scan(&a, &b)
	fmt.Println("amalni kiriting: \n1-qo'shish, \n2-ayirish, \n3-bo'lish, \n4-ko'paytirish")
	fmt.Scan(&amal)
 	switch amal{
 		case 1:
  			fmt.Println(a + b)
 		case 2: 
  			fmt.Println(a - b)
 		case 3: 
  			fmt.Println(a / b)
		case 4: 
			fmt.Println(a * b)
  		default:
			fmt.Println("You can input between 1 and 4!")  
 	}
}