package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	fmt.Printf("Please enter file name :")
	var fileName string
	names := make([]name, 0)
	fmt.Scanf("%s", &fileName)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error open file")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var counter int
	for scanner.Scan() {
		counter++
		line := scanner.Text()
		split := strings.Split(line, " ")
		if len(split) != 2 {
			fmt.Println("Input file error line", counter, "format incorrect")
			continue
		}
		names = append(names, name{
			fname: split[0],
			lname: split[1],
		})

	}

	for _, n := range names {
		fmt.Println(n.fname, n.lname)
	}
}
