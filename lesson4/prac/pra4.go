package main

import "fmt"

func main()  {
	task(12, 6)   
}
func task(a, b int)  {
	if a < b {
	    for i := a; i < b; i++ {
		fmt.Print(i, " ")
	    }
    }else if a > b {
		for j := a; j >= b; j-- {
			fmt.Print(j, " ")
		}
	}
	fmt.Println()
}