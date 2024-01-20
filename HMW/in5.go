package main

import(
	"fmt"
	"math"
)
func main(){
	var r float64
	fmt.Println("radiusni kiriting:")
	fmt.Scan(&r)

	v := 4/3 * math.Pi * math.Pow(r, 3)
	s := 4 * math.Pi * math.Pow(r, 2)
	
	fmt.Println("Hajmi =", v)
	fmt.Println("Yuzi =", s)
}