package main

import (
 "fmt"
 "math"
)

func nine(x float64, n int) float64 {
 return math.Pow(x, float64(n))
}

func main() {
 result1 := nine(1.1, 2)
 fmt.Println("Result for (1.1, 2):", result1)

 result2 := nine(0.5, 4)
 fmt.Println("Result for (0.5, 4):", result2)
}