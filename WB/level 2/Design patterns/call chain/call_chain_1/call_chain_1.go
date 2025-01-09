package main

import "fmt"

type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request string)
}

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) Handler {
	b.next = handler
	return handler
}

func (b *BaseHandler) Handle(request string) {
	if b.next != nil {
		b.next.Handle(request)
	}
}

type LevelOneHandler struct {
	BaseHandler
}

func (o *LevelOneHandler) Handle(request string) {
	if request == "Сброс пароля" {
		fmt.Println("Ваш запрос на сброс пароля принят")
	}else{
		fmt.Println("Ваш запрос передан на другой уровень")
		o.BaseHandler.Handle(request)
	}
}

type LevelTwoHandler struct{
	BaseHandler
}

func(t *LevelTwoHandler) Handle(request string){
	if request == "Проблема с программным обеспечением"{
		fmt.Println("Ваш запрос \"Проблема с программным обеспечением\" принят")
	}else{
		fmt.Println("Ваш запрос передан на другой уровень")
		t.BaseHandler.Handle(request)
	}
}

type LevelThreeHandler struct{
	BaseHandler
}

func(t *LevelThreeHandler) Handle(request string){
	if request == "Аппаратный сбой" {
		fmt.Println("Ваша запрос \"Аппаратный сбой\" был принят")
	}else{
		fmt.Println("Заявка не может быть обработана")
	}
}

func main(){
	levelOne := &LevelOneHandler{}
	levelTwo := &LevelTwoHandler{}
	levelThree := &LevelThreeHandler{}

	levelOne.SetNext(levelTwo).SetNext(levelThree)

	levelOne.Handle("Аппаратный сбой")
	
	levelOne.Handle("Сброс пароля")

	levelThree.Handle("Неизвестна проблема")
}