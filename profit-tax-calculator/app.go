package main

import "fmt"

func main(){
	var revenue, expenses, taxRate float64

	fmt.Print("Revnue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := (revenue - expenses)

	profit := ebt * (1 - (taxRate / 100))

	println("EBT: ", ebt)
	println("Profit: ", profit)
}