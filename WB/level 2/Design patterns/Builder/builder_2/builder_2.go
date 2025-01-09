package main

import "fmt"

type User struct{
	name string
	age int
	email string
	address string
}

type NewUser struct{
	user User
}

func CreateNewUser() *NewUser{
	return &NewUser{}
}

type UserBuilder interface{
	AddName(name string) UserBuilder
	AddAge(age int) UserBuilder
	AddEmail(email string) UserBuilder
	AddAddress(address string) UserBuilder
	CreateUser() User
}

func (n *NewUser) AddName(name string) UserBuilder{
	n.user.name = name
	return n
}

func (n *NewUser) AddAge(age int) UserBuilder{
	n.user.age = age
	return n
}

func (n *NewUser) AddEmail(email string) UserBuilder{
	n.user.email = email
	return n
}

func (n *NewUser) AddAddress(address string) UserBuilder{
	n.user.address = address
	return n
}

func (n *NewUser) CreateUser() User{
	return n.user
}

func main(){
	user := CreateNewUser()
	newUser:= user.AddName("Sasha").
	AddAge(20).
	AddEmail("POP").
	AddAddress("))))").
	CreateUser()

	fmt.Println(newUser)
}