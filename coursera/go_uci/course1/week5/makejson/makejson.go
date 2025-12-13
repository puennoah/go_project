package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	var inputName, inputAddress string
	fmt.Println("Please enter name and address")
	fmt.Scanln(&inputName, &inputAddress)
	//fmt.Scanln(&inputAddress)
	mapJson := map[string]string {
		"name" : inputName,
		"address" : inputAddress,
	}

	outJson, err := json.Marshal(mapJson)
	if err != nil {
		fmt.Println("Error and conversion")
		return
	}
	fmt.Println(string(outJson))
}