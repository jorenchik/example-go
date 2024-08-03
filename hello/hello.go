package main

import ( 
	"fmt"
	"github.com/jorenchik/example-go/greetings"
)

func main() {
	message := greetings.Hello("Someone")
	fmt.Println(message)
}
