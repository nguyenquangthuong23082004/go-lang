package main

import "fmt"

func main() {
	revenue := getUserInput("Nhập doanh thu (Revenue): ")
	expenses := getUserInput("Nhập chi phí (Expenses): ")
	taxRate := getUserInput("Nhập thuế suất (%): ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("Lợi nhuận trước thuế (EBT): %.1f\n", ebt)
	fmt.Printf("Lợi nhuận sau thuế (Profit): %.1f\n", profit)
	fmt.Printf("Tỷ số EBT / Profit (Ratio): %.3f\n", ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
