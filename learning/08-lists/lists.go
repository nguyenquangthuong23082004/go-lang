package main

import "fmt"

func main() {
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println("Original Array:", prices)

	// Lấy phần tử theo chỉ mục
	fmt.Println("Index 2:", prices[2])

	var productNames [4]string
	productNames = [4]string{"A Book"}
	productNames[2] = "A Carpet"
	fmt.Println("Product Names:", productNames)

	// Tạo slice từ vị trí chỉ mục 1 đến hết
	featuredPrices := prices[1:] // [9.99, 45.99, 20.0]
	fmt.Println("Featured Prices (slice):", featuredPrices)

	// Tạo slice từ một slice khác
	highlightedPrices := featuredPrices[:1] // [9.99]
	fmt.Println("Highlighted Prices (slice of slice):", highlightedPrices)

	// Trích xuất bỏ chỉ mục bắt đầu
	fmt.Println("Prices [:3]:", prices[:3]) // [10.99, 9.99, 45.99]

	// Trích xuất bỏ chỉ mục kết thúc
	fmt.Println("Prices [2:]:", prices[2:]) // [45.99, 20.0]
}
