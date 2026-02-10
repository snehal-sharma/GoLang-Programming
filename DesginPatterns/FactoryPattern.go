//Factory method is a creational design pattern which solves the problem of creating product objects without specifying their concrete classes.

package main

import "fmt"

type Payment interface {
	Pay(float64)
}

type CreditCard struct{}

func (cc CreditCard) Pay(amount float64) {
	fmt.Println("Payment through Credit Card :", amount)
}

type UPI struct{}

func (upi UPI) Pay(amount float64) {
	fmt.Println("Payment through UPI :", amount)
}

func PaymentFactory(mode string) Payment {
	switch mode {
	case "CreditCard":
		return CreditCard{}
	case "UPI":
		return UPI{}
	default:
		fmt.Println("Unauthorized payment mode")
		return nil
	}
}

func main() {
	mode := PaymentFactory("UPI")
	mode.Pay(64)
}
