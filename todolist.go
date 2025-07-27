
package main

import (
	"bufio" // For more robust line reading
	"fmt"
	"os"      // For os.Stdin
	"strconv" // Needed for converting integer to string
	"strings" // For trimming whitespace from user input
)

// Global map to hold the to-do lists
var myTodoLists = make(map[string]map[string]string)
var registeredUsers = make(map[string]string)

// register handles user registration
// It directly modifies the global 'registeredUsers' map, so it doesn't need to return it.
func register() { // No parameter needed if accessing global
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n--- Welcome to the Registration Page ---")
	fmt.Print("Enter Your Username: ")
	userName, _ := reader.ReadString('\n')
	userName = strings.TrimSpace(userName)

	// Check if user already exists in the global map
	if _, ok := registeredUsers[userName]; ok {
		fmt.Printf("Error: The username '%s' already exists. Please choose a different one.\n", userName)
		return // Exit function early
	}

	fmt.Print("Enter Your Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	// Store the new user in the global map
	registeredUsers[userName] = password
	fmt.Printf("User '%s' registered successfully!\n", userName)
}

// login handles user login
// It checks against the global 'registeredUsers' map.
func login() bool { // No parameter needed if accessing global
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n--- Welcome to the Login Page ---")
	fmt.Print("Enter Your Username: ")
	userName, _ := reader.ReadString('\n')
	userName = strings.TrimSpace(userName)

	// Check if the username exists
	storedPassword, ok := registeredUsers[userName]
	if !ok {
		fmt.Printf("Error: Username '%s' not found. Please register first or enter a correct username.\n", userName)
		return false
	}

	fmt.Print("Enter Your Password: ")
	enteredPassword, _ := reader.ReadString('\n')
	enteredPassword = strings.TrimSpace(enteredPassword)

	// Compare the entered password with the stored password
	if storedPassword != enteredPassword {
		fmt.Println("Error: Incorrect password. Please try again.")
		return false
	}

	fmt.Printf("Welcome, %s! You are logged in.\n", userName)
	return true
}


func main() {
	reader := bufio.NewReader(os.Stdin) // Main reader for the menu loop
	loggedIn := false

	fmt.Println("Welcome to the Go To-Do List Application!")

	for { // Outer loop for login/registration
		if !loggedIn {
			fmt.Println("\n--- Main Menu ---")
			fmt.Println("1. Register")
			fmt.Println("2. Login")
			fmt.Println("3. Exit")
			fmt.Print("Enter your choice (1-3): ")

			choiceStr, _ := reader.ReadString('\n')
			choiceStr = strings.TrimSpace(choiceStr)

			switch choiceStr {
			case "1":
				register()
			case "2":
				if login() {
					loggedIn = true
					fmt.Println("Successfully logged in!")
				} else {
					fmt.Println("Login failed. Please try again.")
				}
			case "3":
				fmt.Println("Exiting application. Goodbye!")
				return // Exit the program
			default:
				fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
			}
		} else { // User is logged in, show to-do list operations
           
            
			fmt.Println("\n--- To-Do List Operations ---")
			fmt.Println("1. Add a new list/tasks")
			fmt.Println("2. Update an existing task")
			fmt.Println("3. Delete a task")
			fmt.Println("4. Show tasks")
			fmt.Println("5. Logout") // Added logout option
			fmt.Print("Enter your choice (1-5): ")

			choiceStr, _ := reader.ReadString('\n')
			choiceStr = strings.TrimSpace(choiceStr)

			switch choiceStr {
			case "1":
				fmt.Print("Enter the name for the new list (e.g., 'Monday', 'Work'): ")
				listName, _ := reader.ReadString('\n')
				listName = strings.TrimSpace(listName)
				addList(&myTodoLists, listName) // Pass the address of the global map

			case "2":
				fmt.Print("Enter the name of the list (e.g., 'Monday', 'Work') where the task is: ")
				listName, _ := reader.ReadString('\n')
				listName = strings.TrimSpace(listName)

				fmt.Print("Enter the key of the task to update (e.g., 'Task1', 'Task2'): ")
				taskKey, _ := reader.ReadString('\n')
				taskKey = strings.TrimSpace(taskKey)

				updeteList(&myTodoLists, listName, taskKey)

			case "3":
				fmt.Print("Enter the name of the list (e.g., 'Monday', 'Work') from which to delete: ")
				listName, _ := reader.ReadString('\n')
				listName = strings.TrimSpace(listName)

				fmt.Print("Enter the key of the task to delete (e.g., 'Task1', 'Task2'): ")
				taskKey, _ := reader.ReadString('\n')
				taskKey = strings.TrimSpace(taskKey)

				deletList(&myTodoLists, listName, taskKey)

			case "4":
				fmt.Print("Enter the name of the list (e.g., 'Monday', 'Work') to show tasks for: ")
				listName, _ := reader.ReadString('\n')
				listName = strings.TrimSpace(listName)

				fmt.Print("Do you want to see a specific task? (yes/no): ")
				specificChoice, _ := reader.ReadString('\n')
				specificChoice = strings.TrimSpace(specificChoice)

				if strings.ToLower(specificChoice) == "yes" {
					fmt.Print("Enter the key of the specific task (e.g., 'Task1', 'Task2'): ")
					taskKey, _ := reader.ReadString('\n')
					taskKey = strings.TrimSpace(taskKey)
					showTaskList(myTodoLists, listName, taskKey)
				} else {
					showTaskList(myTodoLists, listName)
				}
			case "5": // Logout option
				loggedIn = false
				fmt.Println("Logged out successfully.")
			default:
				fmt.Println("Invalid choice. Please enter a number between 1 and 5.")
			}
		}
	}
}
// addList function (unchanged)
func addList(list *map[string]map[string]string, value string) *map[string]map[string]string {
	reader := bufio.NewReader(os.Stdin) // Local reader for this function

	if *list == nil {
		*list = make(map[string]map[string]string)
	}

	if _, ok := (*list)[value]; !ok {
		(*list)[value] = make(map[string]string)
	}

	contInput := "yes"
	taskCounter := 1

	for contInput != "no" {
		tasks := ""
		fmt.Println("Enter your task, please:")
		tasks, _ = reader.ReadString('\n')
		tasks = strings.TrimSpace(tasks)

		(*list)[value]["Task"+strconv.Itoa(taskCounter)] = tasks
		taskCounter++

		fmt.Printf("Do you want to continue adding tasks to your '%s' list? Enter 'yes' or 'no': ", value)
		contInput, _ = reader.ReadString('\n')
		contInput = strings.TrimSpace(contInput)
	}

	fmt.Println("Your to-do list for", value, ":", *list)
	return list
}

// updeteList function (renamed for consistency and fixed fmt.Printf format)
func updeteList(list *map[string]map[string]string, value string, taskKey string) *map[string]map[string]string {
	reader := bufio.NewReader(os.Stdin) // Local reader for this function

	if *list == nil {
		fmt.Println("Error: The main todo list is nil. Please add tasks first.")
		return list
	}

	dayTasks, ok := (*list)[value]
	if !ok {
		fmt.Printf("Error: The list for '%s' does not exist.\n", value)
		return list
	}

	currentTask, ok := dayTasks[taskKey]
	if !ok {
		fmt.Printf("Error: The task '%s' does not exist in the '%s' list.\n", taskKey, value)
		return list
	}

	newTask := ""
	// Corrected Printf format: remove the extra argument for the string literal
	fmt.Printf("Enter the new task for '%s' in '%s' (it is currently: '%s'): ", taskKey, value, currentTask)
	newTask, _ = reader.ReadString('\n')
	newTask = strings.TrimSpace(newTask)

	dayTasks[taskKey] = newTask // Update the task
	fmt.Println("Updated to-do list:", *list)
	return list
}

// deletList function (unchanged)
func deletList(list *map[string]map[string]string, value string, taskKey string) *map[string]map[string]string {
	if *list == nil {
		fmt.Println("Error: The main todo list is nil. Please add tasks first.")
		return list
	}

	dayTasks, ok := (*list)[value]
	if !ok {
		fmt.Printf("Error: The list for '%s' does not exist.\n", value)
		return list
	}

	if _, ok := dayTasks[taskKey]; !ok {
		fmt.Printf("Error: The task '%s' does not exist in the '%s' list.\n", taskKey, value)
		return list
	}

	delete(dayTasks, taskKey) // Delete the task
	fmt.Println("Updated to-do list after deletion:", *list)
	return list
}

// showTaskList function (unchanged)
func showTaskList(list map[string]map[string]string, value string, taskKey ...string) {
	// 1. Initial checks for the outer map (list)
	if list == nil {
		fmt.Println("Error: The main todo list is nil (uninitialized).")
		return
	}

	// 2. Check if the specified 'day' (value) exists in the outer map
	dayTasks, ok := list[value]
	if !ok {
		fmt.Printf("Error: The list for '%s' does not exist.\n", value)
		return
	}

	// At this point, 'dayTasks' is the inner map for the given 'value' (day)
	// It might be empty, but it's guaranteed not to be nil.

	// 3. Determine if a specific task key was provided (variadic parameter)
	if len(taskKey) > 0 {
		// Scenario 2: User wants to see a specific task for the day
		specificTaskKey := taskKey[0] // Get the first (and only) task key provided

		taskValue, found := dayTasks[specificTaskKey]
		if !found {
			fmt.Printf("Error: Task '%s' not found in the '%s' list.\n", specificTaskKey, value)
			return
		}
		fmt.Printf("Task '%s' on %s: %s\n", specificTaskKey, value, taskValue)
	} else {
		// Scenario 1: User wants to see all tasks for the specified day
		fmt.Printf("All tasks for %s:\n", value)

		if len(dayTasks) == 0 {
			fmt.Printf("  No tasks found for %s.\n", value)
			return
		}

		// Iterate and print all tasks for this day
		for key, task := range dayTasks {
			fmt.Printf("  - %s: %s\n", key, task)
		}
	}
}