package main

import "fmt"

func main(){
	var name string 
	var age int
	var work string 
	fmt.Println("your name")
	fmt.Scanln(&name)
	fmt.Println("your age")
	_, err := fmt.Scanln(&age)
	fmt.Println(err)
	fmt.Println("where do you work")
	fmt.Scanln(&work)
	awe := fmt.Sprintf("my name is %s\n my age is %d\n I work in %s\n", name, age, work)

	fmt.Println(awe)
}


