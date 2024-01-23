package main

import "fmt"

func main(){
	res := str("Hello")
	fmt.Println(res)
}
func str(txt string) string{
	var a string

	for i := len(txt)-1; i >= 0; i-- {
		a += string(txt[i])
	}
	return a
}