package main

import "fmt"

func main() {
	res := thirdd(2, 3)
	fmt.Println(res)
	
}

func thirdd(a, b int) float64  {
	var new = float64(a) / float64(b)
	return new
}