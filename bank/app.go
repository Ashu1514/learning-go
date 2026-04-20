package main

import "fmt"

func main() {
	var accountBalance = 1000.0
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	for {
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int

		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your Balance is", accountBalance)
		} else if choice == 2 {
			fmt.Println("How much do you want to deposit?")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid Amount! Must be greater than 0")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New Amount:", accountBalance)
		} else if choice == 3 {
			fmt.Println("How much do you want to withdraw?")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid Amount! Must be greater than 0")
				continue
			} 
			
			if withdrawalAmount > accountBalance {
				fmt.Println("Insufficiant Balance!")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New Amount:", accountBalance)
		} else {
			fmt.Println("Goodbye!")
			break
		}
		continue
	}
	

}