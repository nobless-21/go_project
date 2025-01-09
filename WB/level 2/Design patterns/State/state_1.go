package main

import "fmt"

type State interface {
	SelectDrink(machine *VendingMachine)
	InsertMoney(machine *VendingMachine)
	DispenseDrink(machine *VendingMachine)
	Cancel(machine *VendingMachine)
}

type WaitingForSelectionState struct{}

func (w *WaitingForSelectionState) SelectDrink(machine *VendingMachine) {
	fmt.Println("Выбирети напиток")
	machine.SetState(&DrinkSelectedState{})
}

func (w *WaitingForSelectionState) InsertMoney(machine *VendingMachine){
	fmt.Println("Сначала выбирите напиток")
}

func (w *WaitingForSelectionState) DispenseDrink(machine *VendingMachine){
	fmt.Println("Сначала выбирите напиток")
}

func (w *WaitingForSelectionState) Cancel(machine *VendingMachine){
	fmt.Println("Отменить операцию невозможно так как напито ещё не выбран")
}

type DrinkSelectedState struct{}

func (d *DrinkSelectedState) SelectDrink(machine *VendingMachine) {
	fmt.Println("Напиток уже выбран")
}

func (d *DrinkSelectedState) InsertMoney(machine *VendingMachine){
	fmt.Println("Проверка наличия напитка")
	machine.SetState(&WaitingForPaymentState{})
}

func (d *DrinkSelectedState) DispenseDrink(machine *VendingMachine){
	fmt.Println("Проверка наличия напитка")
}

func (d *DrinkSelectedState) Cancel(machine *VendingMachine){
	fmt.Println("Отмена операции произошла усешна")
	machine.SetState(&WaitingForSelectionState{})
}

type WaitingForPaymentState struct{}

func (d *WaitingForPaymentState) SelectDrink(machine *VendingMachine) {
	fmt.Println("Напиток уже выбран")
}

func (d *WaitingForPaymentState) InsertMoney(machine *VendingMachine){
	fmt.Println("Напиток есть в наличии")
}

func (d *WaitingForPaymentState) DispenseDrink(machine *VendingMachine){
	fmt.Println("Ожидание оплаты")
	machine.SetState(&DispensingState{})
}

func (d *WaitingForPaymentState) Cancel(machine *VendingMachine){
	fmt.Println("Отмена операции произошла усешна")
	machine.SetState(&WaitingForSelectionState{})
}

type DispensingState struct{}

func (d *DispensingState) SelectDrink(machine *VendingMachine) {
	fmt.Println("Напиток уже выбран")
	machine.SetState(&WaitingForSelectionState{})
}

func (d *DispensingState) InsertMoney(machine *VendingMachine){
	fmt.Println("Напиток есть в наличии")
}

func (d *DispensingState) DispenseDrink(machine *VendingMachine){
	fmt.Println("Оплата произведена")
}

func (d *DispensingState) Cancel(machine *VendingMachine){
	fmt.Println("Отмена операции произошла усешна")
	machine.SetState(&WaitingForSelectionState{})
}

type InsufficientFundsState struct{}

func (d *InsufficientFundsState) SelectDrink(machine *VendingMachine) {
	fmt.Println("Напиток уже выбран")
}

func (d *InsufficientFundsState) InsertMoney(machine *VendingMachine){
	fmt.Println("Напиток есть в наличии")
}

func (d *InsufficientFundsState) DispenseDrink(machine *VendingMachine){
	fmt.Println("Оплата недостаточно средств")
}

func (d *InsufficientFundsState) Cancel(machine *VendingMachine){
	fmt.Println("Отмена операции невозможна")
	machine.SetState(&WaitingForSelectionState{})
}

type VendingMachine struct{
	vendingMachine State
}

func (d *VendingMachine) SelectDrink() {
	d.vendingMachine.SelectDrink(d)
}

func (d *VendingMachine) InsertMoney(){
	d.vendingMachine.InsertMoney(d)
}

func (d *VendingMachine) DispenseDrink(){
	d.vendingMachine.DispenseDrink(d)
}

func (d *VendingMachine) Cancel(){
	d.vendingMachine.Cancel(d)
}

func (d *VendingMachine) SetState(machine State){
	d.vendingMachine = machine
}

func main() {
	vendingMachine := &VendingMachine{vendingMachine: &WaitingForSelectionState{}}

	vendingMachine.SelectDrink()

	vendingMachine.InsertMoney()

	vendingMachine.DispenseDrink()

	vendingMachine.Cancel() 


	vendingMachine.SelectDrink() 
	vendingMachine.InsertMoney()  
	vendingMachine.DispenseDrink()
	vendingMachine.Cancel() 

	vendingMachine.SelectDrink()
	vendingMachine.InsertMoney()
	vendingMachine.DispenseDrink()
}