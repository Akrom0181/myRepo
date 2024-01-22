package main 

import "fmt"

func main(){
	var a string
	fmt.Println("input programming languange:")
	fmt.Scan(&a)

	switch a {
	case "Go":
		fmt.Println("Hello, Go Developer!")
	case "Java":
		fmt.Println("Hello, Java Developer!")
	case "Js":
		fmt.Println("Hello, Js Developer!")
	case "Dart":
		fmt.Println("Hello, Dart Developer!")	
	}
}