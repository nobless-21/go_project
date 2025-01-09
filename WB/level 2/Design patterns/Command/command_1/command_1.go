package main

import "fmt"

type Command interface {
	Execute()
}

type Light struct{}

type Fan struct{}

type Thermostat struct {}

func (n *Light) On() {
	fmt.Println("Свет включён")
}

func (f *Light) Off(){
	fmt.Println("Свет выключен")
}

func (n *Fan) On(){
	fmt.Println("Вентилятор включен")
}

func (f *Fan) Off(){
	fmt.Println("Вентилятор выключен")
}

func (t *Thermostat) SetTemp(temp int){
	fmt.Printf("Температура установлена на %d градусов\n", temp)
}

type LightOnCommand struct{
	light *Light
}

func (l *LightOnCommand) Execute(){
	l.light.On()
}

type LightOffCommand struct{
	light *Light
}

func (l *LightOffCommand) Execute(){
	l.light.Off()
}

type FanOnCommand struct{
	fan *Fan
}

func (f *FanOnCommand) Execute(){
	f.fan.On()
}

type FanOffCommand struct{
	fan *Fan
}

func(f *FanOffCommand) Execute(){
	f.fan.Off()
}

type ThermostatSetCommand struct{
	thermostat *Thermostat
	temp int
}

func (t *ThermostatSetCommand) Execute(){
	t.thermostat.SetTemp(t.temp)
}

type RemoteControl struct{
	command Command
}

func (r *RemoteControl) SetCommand(command Command){
	r.command = command
}

func (r *RemoteControl) PressButton(){
	r.command.Execute()
}

func main(){
	light := &Light{}
	fan := &Fan{}
	thermostat := &Thermostat{}

	lightOn := &LightOnCommand{light: light}
	lightOff := &LightOffCommand{light: light}
	fanOn := &FanOnCommand{fan: fan}
	fanOff := &FanOffCommand{fan: fan}
	setTemp := &ThermostatSetCommand{thermostat: thermostat, temp: 22}

	remote := &RemoteControl{}

	remote.SetCommand(lightOn)
	remote.PressButton()

	remote.SetCommand(fanOn)
	remote.PressButton()

	remote.SetCommand(setTemp)
	remote.PressButton()

	remote.SetCommand(lightOff)
	remote.PressButton()

	remote.SetCommand(fanOff)
	remote.PressButton()
}