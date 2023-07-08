package main

import (
	"fmt"
	"hayashi/module-greetings/greet"
	"log"
)

func main() {
	log.SetPrefix("Greeting: ")
	log.SetFlags(0)

	message, err := greet.Hello("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
