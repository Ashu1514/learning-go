package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main(){
	var investmentAmount, years float64
	expectedReturnRate := 7.5
	// var years float64 = 10

	outputText("Enter a Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Enter a Investment Years: ")
	fmt.Scan(&years)

	outputText("Enter a Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFuturevalue(investmentAmount, expectedReturnRate, years)

	fmt.Printf("Future value is %.2f\n",futureValue)
	fmt.Printf("Future value with inflation %.2f\n", futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFuturevalue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1 + (expectedReturnRate / 100), years)
	rfv = fv / math.Pow(1 + (inflationRate / 100), years)

	return fv, rfv
}