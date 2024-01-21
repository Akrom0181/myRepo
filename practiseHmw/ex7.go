package main

import( 
	"fmt",
   	"math"
)

func main(){
	var r float32 = 5
	var s = math.Pi * r * r
	var l = 2 * math.Pi * r
	fmt.Println("l =", l,"s =", s)
}