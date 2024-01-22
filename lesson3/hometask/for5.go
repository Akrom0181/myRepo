package main 

import(
	"fmt"
    "strconv"
)
func main(){

	a := 1020
	fmt.Printf("a %d\n", a)
	str(a)
}

func str(num int) {
	numStr := strconv.Itoa(num)
	for i := 0; i < len(numStr); i++ {
		fmt.Println("%c ", numStr[i])
	}
	fmt.Println()
}