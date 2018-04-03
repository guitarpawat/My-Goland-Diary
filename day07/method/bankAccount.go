package main

import (
	"fmt"
)

type bankAccount struct {
	name    string
	balance int
}

func (ac bankAccount) toString() string {
	// Concat string and int.
	return fmt.Sprintf("%s: $%d", ac.name, ac.balance)
}

// Need to pass by pointer, so it can be edit.
func (ac *bankAccount) deposit(amount int) bool {
	if amount <= 0 {
		return false
	}
	ac.balance += amount
	return true
}

func (ac *bankAccount) withdraw(amount int) bool {
	if amount <= 0 || amount > ac.balance {
		return false
	}
	ac.balance -= amount
	return true
}

func main() {
	b := bankAccount{"Pawat", 200}
	fmt.Println(b.toString())
	if b.withdraw(300) {
		fmt.Println("New balance:", b.balance)
	} else {
		fmt.Println("Cannot withdraw")
	}

	if b.deposit(100) {
		fmt.Println("New balance:", b.balance)
	} else {
		fmt.Println("Cannot deposite")
	}

	if b.deposit(-100) {
		fmt.Println("New balance:", b.balance)
	} else {
		fmt.Println("Cannot deposite")
	}

	if b.withdraw(300) {
		fmt.Println("New balance:", b.balance)
	} else {
		fmt.Println("Cannot withdraw")
	}
}
