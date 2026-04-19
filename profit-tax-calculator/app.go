package main

import "fmt"

func main(){
	revenue := getUserInput("Revnue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateProfit(revenue, expenses, taxRate)

	fmt.Println("EBT: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Printf("Ratio: %.2f", ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateProfit(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = (revenue - expenses)
	profit = ebt * (1 - (taxRate / 100))	
	ratio = ebt / profit
	return ebt, profit, ratio
}