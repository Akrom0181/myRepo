package main

import "fmt"

func main(){
fun("1241")
}

func fun(a string) {
	
	for i := len(a)-1; i >= 0; i-- {
		fmt.Print(string(a[i]))
	}

}