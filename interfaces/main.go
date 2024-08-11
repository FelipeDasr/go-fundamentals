package main

type IBank interface {
	deposit(amount float64)
	withdraw(amount float64)
}

type Bank1 struct {
	name      string
	propertyX float64
}

func (bank *Bank1) deposit(amount float64) {
	bank.propertyX += amount
}

func (bank *Bank1) withdraw(amount float64) {
	bank.propertyX -= amount
}

func makeOperations(bank IBank, amount float64) {
	bank.deposit(amount)
	bank.withdraw(amount)
}

func main() {
	bank1 := Bank1{name: "Bank 1", propertyX: 100}
	makeOperations(&bank1, 50)
}