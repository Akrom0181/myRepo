package main

import( 
	"fmt"
	"strconv"
)

func main() {

	res := add(1234)
	if res % 2 == 0{
		fmt.Println(true)
	}else {
		fmt.Println(false)
	}
	
}

func add(a int) int {
	var sum = 0
	var str = strconv.Itoa(a)
	for i := 0; i < len(str); i++ {
		a %= 10
		sum += a
	}
	return sum
}