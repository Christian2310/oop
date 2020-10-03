package main

import "fmt"

// Paymethod is an interface to Pay method
type Paymethod interface {
	Pay()
}

// Paypal is an strjuct fro a payment
type Paypal struct{}

// Pay is a method of the interface Paymethod
func (p Paypal) Pay() {
	fmt.Println("Pagado por paypal")
}

// ApplePay is an strjuct fro a payment
type ApplePay struct{}

// Pay is a method of the interface Paymethod
func (a ApplePay) Pay() {
	fmt.Println("Pagado por ApplePay")
}

// Cash is an strjuct fro a payment
type Cash struct{}

// Pay is a method of the interface Paymethod
func (c Cash) Pay() {
	fmt.Println("Pagado con Cash")
}

// Factory is a function to select the payment method
func Factory(method uint) Paymethod {
	switch method {
	case 1:
		return Paypal{}
	case 2:
		return ApplePay{}
	case 3:
		return Cash{}
	default:
		return nil
	}
}

func main() {
	var method uint
	fmt.Println("Select the payment method")
	fmt.Println("\t 1:Paypal 2:ApplePay 3:Cash")
	_, err := fmt.Scanln(&method)
	if err != nil {
		panic("try a valid method")
	}
	if method > 3 {
		panic("try a valid method")
	}

	payMethod := Factory(method)
	payMethod.Pay()

}
