package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Nhập doanh thu (Revenue): ")
	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Nhập chi phí (Expenses): ")
	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Nhập thuế suất (%): ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("Lợi nhuận trước thuế (EBT): %.1f\n", ebt)
	fmt.Printf("Lợi nhuận sau thuế (Profit): %.1f\n", profit)
	fmt.Printf("Tỷ số EBT / Profit (Ratio): %.3f\n", ratio)

	storeResults(ebt, profit, ratio)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Invalid value. Value must be greater than 0.")
	}

	return userInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func storeResults(ebt, profit, ratio float64) {
	resultsText := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(resultsText), 0644)
}
