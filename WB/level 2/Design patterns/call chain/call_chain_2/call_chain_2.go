package main

import "fmt"

type OrderHandler interface {
	SetNext(handler OrderHandler) OrderHandler
	Handle(orderSatus string)
}

type OrderManager struct {
	next OrderHandler
}

func (o *OrderManager) SetNext(handler OrderHandler) OrderHandler {
	o.next = handler
	return handler
}

func (o *OrderManager) Handle(orderStatus string) {
	if o.next != nil {
		o.next.Handle(orderStatus)
	}
}

type NewOrderHandler struct {
	OrderManager
}

func (n *NewOrderHandler) Handle(orderStatus string) {
	if orderStatus == "Новый заказ" {
		fmt.Println("Заказ только что создан и требует подтверждения")
	}else{
		n.OrderManager.Handle(orderStatus)
	}
}

type ConfirmedOrderHandler struct{
	OrderManager
}

func  (c *ConfirmedOrderHandler) Handle(orderStatus string){
	if orderStatus == "Подтвержденный заказ" {
		fmt.Println("Заказ подтвержден и готов к обработке")
	}else{
		c.OrderManager.Handle(orderStatus)
	}
}

type ShippedOrderHandler struct{
	OrderManager
}

func (s *ShippedOrderHandler) Handle(orderStatus string){
	if orderStatus == "Отправленный заказ" {
		fmt.Println("Заказ отправлен и ожидает доставки")
	}else{
		s.OrderManager.Handle(orderStatus)
	}
}

type DeliveredOrderHandler struct{
	OrderManager
}

func (d *DeliveredOrderHandler) Handle(orderStatus string){
	if orderStatus == "Доставленный заказ" {
		fmt.Println("Заказ доставлен клиенту")
	}else{
		fmt.Println("Cтатус заказа не соответствует ни одному из обработчиков")
	}
}

func main(){
	newOrder := &NewOrderHandler{}
	Confirmed := &ConfirmedOrderHandler{}
	Shipped := &ShippedOrderHandler{}
	Delivered := &DeliveredOrderHandler{}

	newOrder.SetNext(Confirmed).SetNext(Shipped).SetNext(Delivered)

	newOrder.Handle("Новый заказ")
	newOrder.Handle("Подтвержденный заказ")
	newOrder.Handle("Отправленный заказ")
	newOrder.Handle("Доставленный заказ")
	newOrder.Handle("Неизвестный статус")
}