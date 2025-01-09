package main

import "fmt"

type Command interface {
	Execute()
	Undo()
}

type Task struct {
	name   string
	status bool
}

type AddTaskCommand struct{
	task *Task
	name string
}

func (a *AddTaskCommand) Execute() {
	a.task.name = a.name
	fmt.Printf("Команда %s добавлена\n", a.name)
}

func (a *AddTaskCommand) Undo(){
	a.task.name = " "
	fmt.Printf("Команда %s была отменена\n", a.name)
}

type RemoveTaskCommand struct{
	task *Task
	name string
}

func (r *RemoveTaskCommand) Execute(){
	r.task.name = " "
	r.task.status = false
	fmt.Printf("Команда %s удалена\n", r.name)
}

func (r *RemoveTaskCommand) Undo(){
	r.task.name = r.name
	r.task.status = false
	fmt.Printf("Команда %s была востановлена\n", r.name)
}

type CompleteTaskCommand struct{
	task *Task 
}

func (c *CompleteTaskCommand) Execute(){
	c.task.status = true
	fmt.Printf("Команда %s выполнена\n", c.task.name)
}

func (c *CompleteTaskCommand) Undo(){
	c.task.status = false
	fmt.Printf("Команда %s была отменена\n", c.task.name)
}

type TaskManager struct{
	command Command
}

func(t *TaskManager) SetCommand(command Command){
	t.command = command
}

func(t *TaskManager) PressButton(){
	t.command.Execute()
}

func(t *TaskManager) Undo(){
	t.command.Undo()
}

func main() {
	task := &Task{}
	manager := &TaskManager{}

	addCommand := &AddTaskCommand{task: task, name: "Сделать домашнее задание"}
	manager.SetCommand(addCommand)
	manager.PressButton()

	completeCommand := &CompleteTaskCommand{task: task}
	manager.SetCommand(completeCommand)
	manager.PressButton()


	manager.Undo()

	removeCommand := &RemoveTaskCommand{task: task, name: "Сделать домашнее задание"}
	manager.SetCommand(removeCommand)
	manager.PressButton()

	manager.Undo()
}