package main 

import "fmt"

func main(){

	var sek float64
	fmt.Print("Sekund kiriting: ")
	fmt.Scan(&sek)

	convert := sek / 60

	fmt.Println(convert, "minut")


}