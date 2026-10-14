package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("My Todo List")

	var input string
	var todos []string
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	for {
		fmt.Println("Enter a todo: ")

		if !scanner.Scan() {
			break
		}

		input = scanner.Text()

		if input == "" {
			continue
		} else if input == "done" {
			break
		} else {
			todos = append(todos, input)
		}
	}
	fmt.Println("Todos:")

	for i, todo := range todos {
		fmt.Println(i+1, todo)
	}
}
