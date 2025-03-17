package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	taskName  string
	completed bool
}

var tasks []Task

func addTask(task string) {
	newTask := Task{taskName: task, completed: false}
	tasks = append(tasks, newTask)

	fmt.Println("Task Added")
}

func listTask() {
	for i, task := range tasks {
		status := "not yet"
		if task.completed {
			status = "done"
		}

		fmt.Println(i+1, task.taskName, status)
	}

}

func markCompleted(index int) {
	if index >= 1 && index <= len(tasks) {
		tasks[index-1].completed = true
		fmt.Println("Task Markes as Complete")
	} else {
		fmt.Println("invalid index")
	}
}

func updateTask(index int, newString string) {
	if index >= 1 && index <= len(tasks) {
		tasks[index-1].taskName = newString
		fmt.Println("Task Updated succesfully")
	} else {
		fmt.Println("Invalid Index")
	}
}

func deleteTask(index int) {
	if index >= 1 && index <= len(tasks) {
		tasks = append(tasks[:index-1], tasks[index:]...)
		fmt.Println("Task Deleted succesfully")
	} else {
		fmt.Println("Invalid Index")
	}
}

func main() {
	println("task-traker v-0.1.0")
	var indexInput int
	var taskInput, newTaskInput string

	fmt.Println("Options")
	fmt.Println("1. Add Task")
	fmt.Println("2. List Task")
	fmt.Println("3. Mark as Completed")
	fmt.Println("4. Edit Task")
	fmt.Println("5. Delete Task")
	fmt.Println("6. Exit")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter Choice (1 - 6) : ")
		scanner.Scan()
		input := scanner.Text()

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid Choice")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter Text: ")
			scanner.Scan()
			taskInput = scanner.Text()
			addTask(taskInput)
		case 2:
			listTask()
		case 3:
			fmt.Print("Enter Index: ")
			scanner.Scan()

			indexInput, _ = strconv.Atoi(scanner.Text())
			markCompleted(indexInput)
		case 4:
			fmt.Print("Enter Index: ")
			scanner.Scan()
			indexInput, _ = strconv.Atoi(scanner.Text())

			fmt.Print("Enter Text: ")
			scanner.Scan()
			newTaskInput = scanner.Text()
			updateTask(indexInput, newTaskInput)
		case 5:
			fmt.Print("Enter Index: ")
			scanner.Scan()
			indexInput, _ = strconv.Atoi(scanner.Text())
			deleteTask(indexInput)
		case 6:
			os.Exit(0)
		default:
			fmt.Println("invalid choice")
		}
	}
}
