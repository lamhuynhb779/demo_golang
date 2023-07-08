package main

import (
	"fmt"
	"hayashi/greetings/pk_greetings"
	"log"
)

func main() {

	log.SetPrefix("Error: ")
	log.SetFlags(0)

	message, err := pk_greetings.Hello("Lam")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

}
