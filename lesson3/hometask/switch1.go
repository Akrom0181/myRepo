package main

import "fmt"

func main(){
	var a int
	fmt.Println("Input number:")
	fmt.Scan(&a)
	if Tub(a){
		fmt.Println("tub")
	}else {
		fmt.Println("tub emas")
	}
}
func Tub(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}