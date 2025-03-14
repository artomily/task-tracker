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

func listTask(){
	for i, task := range tasks{
		status := "not yet"
		if task.completed {
			status = "done"
		}

		fmt.Println("%d. %s [$s]\n", i+1, task.taskName, status)
	}

}

func markCompleted(index int){
	if index >= 1 && index <= len(tasks){
		tasks[index - 1].completed = true
		fmt.Println("Task Markes as Complete")
	}else{
		fmt.Println("invalid index")
	}
}

