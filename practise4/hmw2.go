package main

import (
 "fmt"
)

type Product struct {
 Name     string
 Quantity int
}

func main() {
 	products := []Product{
  		{Name: "Mirinda", Quantity: 5},
  		{Name: "Fanta", Quantity: 2},
  		{Name: "Sprite", Quantity: 3},
  		{Name: "Pepsi", Quantity: 4},
 	 	{Name: "Coke", Quantity: 10},
    }

maxQuantity := 0
maxProduct := ""

 	for _, product := range products {
 		if product.Quantity > maxQuantity {
   			maxQuantity = product.Quantity
   			maxProduct = product.Name
  		}
 	}

fmt.Printf("the most product is %s with a quantity of %d", maxProduct, maxQuantity)
}