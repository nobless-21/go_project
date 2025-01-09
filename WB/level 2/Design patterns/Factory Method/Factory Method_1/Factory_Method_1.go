package main

import "fmt"

type Vehicle interface {
	Drive() string
	GetCapacity() int
}

type Car struct{}

func (c *Car) Drive() string{
	return "Driving a car."
}

func (c *Car) GetCapacity() int{
	return 5
}

type Truck struct{}

func (t *Truck) Drive() string{
	return "Driving a truck."
}

func (t *Truck) GetCapacity() int{
	return 7
}

type Motorcycle struct{}

func (m *Motorcycle) Drive() string{
	return "Driving a motorcycle."
}

func (m *Motorcycle) GetCapacity() int{
	return 2
}

type VehicleFactory interface {
	CreateVehicle() Vehicle
}

type CarFactory struct{}

func (f *CarFactory) CreateVehicle() Vehicle{
	return &Car{}
}

type TruckFactory struct{}

func(f *TruckFactory) CreateVehicle() Vehicle{
	return &Truck{}
}

type MotorcycleFactory struct{}

func (f *MotorcycleFactory) CreateVehicle() Vehicle{
	return &Motorcycle{}
}

func main(){
	var carFactory VehicleFactory

	carFactory = &CarFactory{}
	car1 := carFactory.CreateVehicle()
	fmt.Println(car1.GetCapacity())
}