package main

import(
	"fmt"
	"math"
)
func main(){
	
	var(
		weight float64
		height float64
	)
fmt.Println("input your weight: ")
fmt.Scan(&weight)

fmt.Println("input your height: ")
fmt.Scan(&height)

bmi := weight / math.Pow(height, 2)
fmt.Print("your BMI is ", bmi)
}