package main

import (
	"fmt"
	"os"
)

const storeCalulationFile = "calculation"

func writeDataToFile(ebt, profit, ratio float64) error {
	dataText := fmt.Sprintf("ebt: %.2f, profit: %.2f, ratio: %.2f", ebt, profit, ratio)
	return os.WriteFile(storeCalulationFile, []byte(dataText), 0644)
}

func main(){
	revenue := getUserInput("Revnue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateProfit(revenue, expenses, taxRate)

	err := writeDataToFile(ebt, profit, ratio)

	if err != nil {
		fmt.Println("Error while writing data to file!", err)
	}

	fmt.Println("EBT: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		panic("Zero or Negative values are not allowed!")
	}
	return userInput
}

func calculateProfit(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = (revenue - expenses)
	profit = ebt * (1 - (taxRate / 100))	
	ratio = ebt / profit
	return ebt, profit, ratio
}