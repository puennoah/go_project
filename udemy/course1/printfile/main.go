package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please input a file name for print out!")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()

}
