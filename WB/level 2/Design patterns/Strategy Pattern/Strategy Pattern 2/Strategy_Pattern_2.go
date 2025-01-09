package main

import "fmt"

type TaxStrategy interface{
	CalculateTax(income float64) float64
}

type SalaryTax struct{}

func (s *SalaryTax) CalculateTax(income float64) float64{
	return income * 0.2
}

type InvestmentTax struct{}

func (i *InvestmentTax) CalculateTax(income float64) float64{
	return income *0.15
}

type BusinessTax struct{}

func(b *BusinessTax) CalculateTax(income float64) float64{
	return income *0.25
}

type TaxCalculator struct{
	tax TaxStrategy
}

func (t *TaxCalculator) SetTypeTax(tax TaxStrategy){
	t.tax = tax
}

func (t *TaxCalculator) Calculate(income float64) float64{
	if t.tax == nil{
		fmt.Println("Установите тип налога")
		return income
	}
	return t.tax.CalculateTax(income)
}

func main(){
	var tax TaxCalculator

	tax.SetTypeTax(&SalaryTax{})
	fmt.Println(tax.Calculate(100.0))

	tax.SetTypeTax(&BusinessTax{})
	fmt.Println(tax.Calculate(100.0))

	tax.SetTypeTax(&InvestmentTax{})
	fmt.Println(tax.Calculate(100.0))
}