package pk_greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("Empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi, %v, wellcome",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
