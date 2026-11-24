package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("My Todo List")

	var input string
	var todos []string
	var number int
	var err error
	scanner := bufio.NewScanner(os.Stdin)

	for { // This is menu
		fmt.Println("===== My Todo List =====")
		fmt.Println("")
		fmt.Println("1. View todos")
		fmt.Println("2. Add todo")
		fmt.Println("3. Delete todo")
		fmt.Println("4. Edit todo")
		fmt.Println("5. Exit")
		fmt.Println("")
		fmt.Println("Choose an option: ")

		if !scanner.Scan() {
			break
		}

		input = scanner.Text()

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
		}
		if input == "" {
			continue
		} else {
			number, err = strconv.Atoi(input)
			if err != nil {
				fmt.Println("Try again with interger.")
				continue
			}

			if number < 1 || number > 5 {
				fmt.Println("Invalid option!")
				continue
			}
		}

		switch number {
		case 1: // View todos
			viewTodos(todos)

		case 2: // Add todo
			todos = addTodo(todos, scanner)
			viewTodos(todos)

		case 3: //Delete todo
			if !viewTodos(todos) {
				break
			}
			for {
				fmt.Println("Which todo number do you want to delete?")

				if !scanner.Scan() {
					break
				}

				input = scanner.Text()
				if err := scanner.Err(); err != nil {
					fmt.Println("Error reading input:", err)
				}
				if input == "" {
					continue
				} else {
					number, err := strconv.Atoi(input)
					if err != nil {
						fmt.Println("Try again with interger.")
						continue
					}

					if number < 1 || number > len(todos) {
						fmt.Println("Invalid todo number")
						continue
					}

					todos = append(todos[:(number-1)], todos[number:]...)
					break

				}
			}
			if err := scanner.Err(); err != nil {
				fmt.Println("Error reading input:", err)
			}
			viewTodos(todos)
		case 4: // Edit todo
			if !viewTodos(todos) {
				break
			}
			for {
				fmt.Println("Which todo number do you want to edit?")

				if !scanner.Scan() {
					break
				}

				input = scanner.Text()
				if input == "" {
					continue
				} else {
					number, err = strconv.Atoi(input)
					if err != nil {
						fmt.Println("Try again with interger.")
						continue
					}

					if number < 1 || number > len(todos) {
						fmt.Println("Invalid todo number")
						continue
					}
				}

				fmt.Println("Enter the new todo:")
				if !scanner.Scan() {
					break
				}

				input = scanner.Text()

				if input == "" {
					continue
				}

				todos[number-1] = input
				break

			}

			if err := scanner.Err(); err != nil {
				fmt.Println("Error reading input:", err)
			}

			viewTodos(todos)

		case 5:
			fmt.Println("Exit")
		}
		if number == 5 {
			break
		}

	}

}

func viewTodos(todos []string) bool {
	if len(todos) < 1 {
		fmt.Println("You got no todo!")
		return false
	}
	fmt.Println("You have", len(todos), "todos")
	fmt.Println("Todos:")
	for i, todo := range todos {
		fmt.Println(i+1, todo)
	}
	return true
}

func addTodo(todos []string, scanner *bufio.Scanner) []string {
	for {
		fmt.Println("Enter a todo: ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		if input == "" {
			continue
		} else if input == "done" {
			break
		} else {
			todos = append(todos, input)
		}
	}
	return todos

}
