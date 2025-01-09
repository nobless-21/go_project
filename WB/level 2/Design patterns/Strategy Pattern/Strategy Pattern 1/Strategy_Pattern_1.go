package main

import "fmt"

type PaymentSrategy interface {
	pay(amount float64)
}

type CreditCardPayment struct{}

func (c *CreditCardPayment) pay(amount float64) {
	fmt.Printf("Оплата будет произведена кредитной картой в размере %f \n", amount)
}

type PayPalPayment struct{}

func (p *PayPalPayment) pay(amount float64){
	fmt.Printf("Оплата будет произведена с помощью PayPal в размере %f \n", amount)
}

type CryptoPayment struct{}

func (c *CryptoPayment) pay(amount float64){
	fmt.Printf("Оплата будет произведена с помощью крипты в размере %f \n", amount)
}

type ShoppingCart struct{
	strategy PaymentSrategy
}

func (s *ShoppingCart) SetStrategy(strategy PaymentSrategy){
	s.strategy = strategy
}

func (s *ShoppingCart) checkout(amount float64){
	if s.strategy == nil {
		fmt.Println("Выберите способ оплаты")
		return
	}
	s.strategy.pay(amount)
}

func main(){
	cart := &ShoppingCart{}

	cart.SetStrategy(&CreditCardPayment{})
	cart.checkout(100.0)

	cart.SetStrategy(&PayPalPayment{})
	cart.checkout(200.0)

	cart.SetStrategy(&CryptoPayment{})
	cart.checkout(300.0)
}