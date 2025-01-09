package main

import (
	"fmt"
)
type Home struct{
	NumberOfRooms int
	garage bool
	garden bool
	type_of_roof string
}

type HomeBuilder interface{
	NumberRooms(NumberOfRooms int) HomeBuilder
	Garage(status bool) HomeBuilder
	Garden(status bool) HomeBuilder
	TypeRoof(typ string) HomeBuilder
	CreateHome() Home
}

type NewHomeBuilder struct{
	home Home
}

func CreateHomeBuilder() *NewHomeBuilder {
	return &NewHomeBuilder{}
}

func (c *NewHomeBuilder) NumberRooms(NumberOfRooms int) HomeBuilder{
	c.home.NumberOfRooms = NumberOfRooms
	return c
}

func (c *NewHomeBuilder) Garage(status bool) HomeBuilder{
	c.home.garage = status
	return c
}

func (c *NewHomeBuilder) Garden(status bool) HomeBuilder{
	c.home.garden = status
	return c
}

func (c *NewHomeBuilder) TypeRoof(typ string) HomeBuilder{
	c.home.type_of_roof = typ
	return c
}

func (c *NewHomeBuilder) CreateHome() Home {
	return c.home
}

func main(){
	createHome := CreateHomeBuilder()
	home := createHome.NumberRooms(2).
		Garage(true).
		Garden(false).
		TypeRoof("ступенчатая").
		CreateHome()

	
		fmt.Printf("Home [Rooms: %d, Garage: %t, Garden: %t, Roof Type: %s]\n", home.NumberOfRooms, home.garage, home.garden, home.type_of_roof)
}