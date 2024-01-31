package main

import "fmt"

type Clientt struct{
	Name  string
	Basket []int
}

func main(){
	clients := []Clientt{
	{Name: "Abbos", Basket: []int{10000, 20000, 5000}},
	{Name: "Ali", Basket: []int{10000, 20000, 5000}},
	{Name: "Davron", Basket: []int{3000,20000, 5000}},
	}
	for _, client := range clients {
		sum := 0
		for _, item := range client.Basket {
			sum += item
		}
		fmt.Printf("%s's basket sum: %dn", client.Name, sum)
	}
}