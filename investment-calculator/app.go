package main

import (
	"fmt"
	"math"
)

func main(){
	const inflationRate float64 = 2.5
	var investmentAmount, years float64
	expectedReturnRate := 7.5
	// var years float64 = 10

	outputText("Enter a Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Enter a Investment Years: ")
	fmt.Scan(&years)

	outputText("Enter a Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	var futureValue = investmentAmount * math.Pow(1 + (expectedReturnRate / 100), years)
	futureRealValue := futureValue / math.Pow(1 + (inflationRate / 100), years)

	fmt.Printf("Future value is %.2f\n",futureValue)
	fmt.Printf("Future value with inflation %.2f\n", futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}