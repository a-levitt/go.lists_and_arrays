package main

import "fmt"

func main() {
	var productNames [4]string
	productNames = [4]string{"A book", "A T-shirt", "A notebook", "A desktop"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices[1])
	fmt.Println(productNames[1])
}