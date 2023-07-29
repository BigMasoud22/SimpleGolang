package main

import (
	"time"
)

func main() {
	person := Person{Name: "Masoud", Age: 18}
	mytask := Task{Name: "Doing the dishes", DueDate: time.Now()}

	person.AddTask(mytask)

	mytask2 := Task{Name: "Cooking dinner !", DueDate: time.Now().Add(2 * time.Hour)}
	person.AddTask(mytask2)
    person.DeleteTask(mytask2)

	person.PrintTasks()
}
