package main

import(
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct{
	taskName string
	completed bool
}

var tasks []Task

func main(){
	println("task-traker v-0.1.0")


}

func addTask(task string){
	newTask := Task{taskName: task, completed: false}
	tasks = append(tasks, newTask)

	fmt.Println("Task Added")
}
