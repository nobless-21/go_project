package main

import "fmt"

type Animal interface {
	Speak() string
	GetType() string
}

type Dog struct{}

func (d *Dog) Speak() string {
	return "Гав!"
}

func (d *Dog) GetType() string {
	return "Собака"
}

type Cat struct{}

func (c *Cat) Speak() string {
	return "Мяу!"
}

func (c *Cat) GetType() string {
	return "Кошка"
}

type Bird struct{}

func (b *Bird) Speak() string {
	return "Чирик!"
}

func (b *Bird) GetType() string {
	return "Птица"
}

type AnimalFactory interface {
	CreateAnimal() Animal
}

type DogFactory struct{}

func (d *DogFactory) CreateAnimal() Animal {
	return &Dog{}
}

type CatFactory struct{}

func (c *CatFactory) CreateAnimal() Animal {
	return &Cat{}
}

type BirdFactory struct{}

func (b *BirdFactory) CreateAnimal() Animal {
	return &Bird{}
}

func main() {
	var animal1, animal2, animal3 AnimalFactory

	animal1 = &DogFactory{}
	dog := animal1.CreateAnimal()
	fmt.Println(dog.Speak())
	fmt.Println(dog.GetType())

	animal2 = &CatFactory{}
	cat := animal2.CreateAnimal()
	fmt.Println(cat.Speak())
	fmt.Println(cat.GetType())

	animal3 = &BirdFactory{}
	bird := animal3.CreateAnimal()
	fmt.Println(bird.Speak())
	fmt.Println(bird.GetType())
}