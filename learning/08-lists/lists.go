package main

import "fmt"

func main() {
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	fmt.Println(prices[2])

	var productNames [4]string
	productNames = [4]string{"A Book"}
	productNames[2] = "A Carpet"
	fmt.Println(productNames)

	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)
}
