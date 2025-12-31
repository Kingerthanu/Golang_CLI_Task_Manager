package main

import (
	"CLI_Task_Manager/task_management"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {

	create := flag.Bool("create", false, "Create a new task")
	list := flag.Bool("list", false, "List all tasks")
	delete := flag.Int("delete", -1, "Delete task by ID")
	name := flag.String("name", "", "Task name")
	desc := flag.String("desc", "", "Task description")
	status := flag.Int("status", 0, "Task status")
	done := flag.Bool("done", false, "Task completion status")
	live := flag.Bool("live", false, "Loop And Update Manager")

	flag.Parse()

	dataDir := "data"
	tasksFile := filepath.Join(dataDir, "salmon.json")

	// Create directory if it doesn't exist
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		log.Fatal("Error creating directory:", err)
	}

	var manager *task_management.TaskManager

	// Try to load existing file
	manager, err := task_management.LoadFromFile(tasksFile)
	if err != nil {
		// File doesn't exist or is invalid, create new manager
		err, manager = task_management.MakeTaskManager("Demo Task")
		if err != nil {
			log.Fatal("Error: Failed Creation Of Task Manager:", err)
		}
		manager.SetPath(tasksFile)

		// Create initial empty file
		if err := manager.SaveToFile(); err != nil {
			log.Fatal("Error creating initial file:", err)
		}
	}

	if *create {

		err := manager.CreateTask(*name, *desc, *status, *done, manager.GetNextID())
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		manager.IncrementID()
		fmt.Println("Task Created Successfully!")

	} else if *list {

		tasks := manager.ListTasks()
		for _, t := range tasks {
			fmt.Printf("ID: %d, Name: %s, Done: %t\n", t.GetID(), t.GetName(), t.IsDone())
		}

	} else if *delete >= 0 {

		err := manager.DeleteTask(*delete)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("Task Deleted Successfully")

	} else if *live {

		defer manager.SaveToFile()
		runInteractiveMode(manager)

	} else {

		flag.Usage()

	}

}

func runInteractiveMode(manager *task_management.TaskManager) {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\n=== Task Manager Interactive Mode ===")
	fmt.Println("Commands: create, list, delete, update, exit")

	for {

		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		command := parts[0]

		switch command {

		case "exit":
			fmt.Println("Exiting...\n")
			return

		case "create":
			fmt.Print("Task Name: ")
			scanner.Scan()
			name := scanner.Text()

			fmt.Print("Description: ")
			scanner.Scan()
			desc := scanner.Text()

			fmt.Print("Status (0): ")
			scanner.Scan()
			statusStr := scanner.Text()
			status := 0
			if statusStr != "" {
				status, _ = strconv.Atoi(statusStr)
			}

			err := manager.CreateTask(name, desc, status, false, manager.GetNextID())
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				manager.IncrementID()
				fmt.Println("Task Created Successfully!")
			}

		case "list":
			tasks := manager.ListTasks()
			if len(tasks) == 0 {
				fmt.Println("No Tasks Found.")
			}
			for _, t := range tasks {
				fmt.Printf("ID: %d | Name: %s | Status: %d | Done: %t\n", t.GetID(), t.GetName(), t.GetStatus(), t.IsDone())
			}

		case "delete":
			if len(parts) < 2 {
				fmt.Println("Usage: delete <id>")
				continue
			}

			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Invalid ID")
				continue
			}
			err = manager.DeleteTask(id)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Task Deleted Successfully!")
			}

		case "update":
			if len(parts) < 2 {
				fmt.Println("Usage: update <id>")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Invalid ID")
				continue
			}

			task := manager.GetTask(id)
			if task != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Print("New name (leave empty to keep current): ")
			scanner.Scan()
			name := scanner.Text()
			if name != "" {
				task.SetName(name)
			}

			fmt.Print("Mark as done? (y/n): ")
			scanner.Scan()
			doneStr := scanner.Text()
			if doneStr == "y" || doneStr == "Y" {
				task.SetDone(true)
			}

			fmt.Println("Task updated successfully!")

		default:
			fmt.Println("Unknown command. Available: create, list, delete, update, exit")
		}

	}

}
