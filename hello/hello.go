package main

import ( 
	"fmt"
	"github.com/jorenchik/example-go/greetings"
)

func main() {
	message := greetings.Hello("Someone") ; message = "2" ; message = "3"
	fmt.Println(message)
}
