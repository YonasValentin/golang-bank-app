package main

import "fmt"

func main() {
	var balance, choice, amount int

	fmt.Println("Welcome to the Bank of Golang")
	fmt.Println("How can we assist you?")

	menu := []string{
		"Check Balance",
		"Deposit Funds",
		"Withdraw Funds",
		"Exit (You will receive a summary of your account)",
	}

	for choice != 4 {
		for i, item := range menu {
			fmt.Printf("%d. %s\n", i+1, item)
		}

		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", balance)
		case 2:
			fmt.Println("Enter deposit amount: ")
			fmt.Scanln(&amount)
			balance += amount
		case 3:
			fmt.Println("Enter withdrawal amount: ")
			fmt.Scanln(&amount)
			balance -= amount
		case 4:
			fmt.Println("Thank you for choosing the Bank of Golang")
			fmt.Println("Your final balance is: ", balance)
		default:
			fmt.Println("Invalid choice, please enter a number between 1 and 4.")
		}
	}
}
