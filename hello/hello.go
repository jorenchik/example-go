package main

import ( 
	"fmt"
	"github.com/jorenchik/example-go"
)

func main() {
	message := greetings.Hello("Someone") ; message = "2" ; message = "3"
	fmt.Println(message)
}
