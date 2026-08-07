// package main

// import "fmt"

// // Task 7: Create a product struct (title, id, price)
// type Product struct {
// 	id    string
// 	title string
// 	price float64
// }

// func main() {
// 	// ----------------------------------------------------------------------
// 	// Task 7 (Bonus): Dynamic array of structs & append a 3rd product
// 	// ----------------------------------------------------------------------
// 	products := []Product{
// 		{
// 			id:    "p1",
// 			title: "Go Programming Book",
// 			price: 29.99,
// 		},
// 		{
// 			id:    "p2",
// 			title: "Ergonomic Keyboard",
// 			price: 99.99,
// 		},
// 	}
// 	fmt.Println("Task 7 - Initial Products:", products)

// 	// ----------------------------------------------------------------------
// 	// Gộp 2 Slice (Unpacking operator ...)
// 	// ----------------------------------------------------------------------
// 	prices := []float64{10.99, 8.99, 5.99}
// 	discountPrices := []float64{101.99, 80.99, 20.59}

// 	// append(prices, discountPrices) // LỖI! Vì discountPrices là []float64 chứ không phải float64
// 	// Sử dụng cú pháp `...` để rã (unpack) các phần tử trong discountPrices
// 	prices = append(prices, discountPrices...)
// 	fmt.Println("Merged prices (với toán tử ...):", prices)
// }