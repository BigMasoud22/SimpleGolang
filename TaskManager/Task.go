package main

import (
	"fmt"
	"time"
)

type Task struct {
	Name    string
	DueDate time.Time
	User    Person
}

type Person struct {
	Name  string
	Age   int
	Tasks []Task
}

func (person *Person) AddTask(task Task) {
	(*person).Tasks = append(person.Tasks, task)
}

func (person *Person) DeleteTask(task Task) bool {

	index, isSuccess := person.FindTask(task)
	if isSuccess {
		(*person).Tasks = append(person.Tasks[:index], person.Tasks[index+1:]...)
		return true
	}
	return false

}
func (person Person) FindTask(task Task) (int, bool) {
	for i, t := range person.Tasks {
		if t.Name == task.Name && t.User.Name == task.User.Name {
			return i, true
		}
	}
	return -1, false
}
func (person Person) PrintTasks() {

	for _, task := range person.Tasks {
		fmt.Println(task.Name)
	}

}
