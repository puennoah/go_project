package main

import (
	"fmt"
)

/*
	#cgo LDFLAGS: -lm
	#include <stdio.h>

	void hello(){
		printf("Hello World\n");
	}
*/
import "C"

func main() {
	fmt.Println("Hi from GO")
	C.hello()
}
