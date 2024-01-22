package main

import "fmt"

func main(){
	var angle int
	fmt.Println("input number:")
	fmt.Scan(&angle)
	switch angle {
	case 3:
		fmt.Println("Uchburchak")
	case 4:
		fmt.Println("to'rtburchak")	
	case 5:
		fmt.Println("beshburchak")
	case 6:
		fmt.Println("oltiburchak")	
	default: 
		fmt.Println("3 dan 6 gacha kiriting!")		
	}
}