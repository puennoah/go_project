package main
import (
	"fmt"
)


func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		mess := "In anonymous function"
		c1 <- mess
		stress := <- c2
		fmt.Println(mess, "AND", stress)
	}()

	mess := "In main() message"
	var stress string	
	select {
	case c2 <- mess:
	case stress = <- c1:
	default: fmt.Println(mess, "AND", stress)
	}
}