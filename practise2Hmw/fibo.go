package main

import "fmt"

func main() {
    n := 10
    a, b := 0, 1
    fmt.Println("Fibonacci series: ")
    for i := 0; i < n; i++ {
        fmt.Print(a, " \n")
        a, b = b, a+b
    }
}
