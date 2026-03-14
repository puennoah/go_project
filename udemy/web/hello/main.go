package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel world"

	fmt.Println(whatToSay)

	i = 13
	fmt.Println("i is set to ", i)

	whatWasSaid ,theOtherThingThatWasSaid := saySomething()
	fmt.Println("The fucntion returned :", whatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "something" , "another something"
}
