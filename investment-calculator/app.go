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

	fmt.Print("Enter a Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter a Investment Years: ")
	fmt.Scan(&years)

	fmt.Print("Enter a Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	var futureValue = investmentAmount * math.Pow(1 + (expectedReturnRate / 100), years)
	futureRealValue := futureValue / math.Pow(1 + (inflationRate / 100), years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}