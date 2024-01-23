// package main

// import(
// 	"fmt"
// 	"strconv"
// )
// func main() {

	
// }

// func five(a int) {
// 	var txt = strconv.Itoa(a)
// 	var b byte
// 	for i := 1; i <= len(txt); i++ {
// 		b += txt[i] 
// 		fmt.Println(b)
// 	}
// }
package main

import (
 "fmt"
 "strconv"
)

func reverseDigits(n int) {
 str := strconv.Itoa(n)
 for i := len(str) - 1; i >= 0; i-- {
  fmt.Print(string(str[i]) + " ")
 }
}

func main() {
 num := 12345
 reverseDigits(num)
}