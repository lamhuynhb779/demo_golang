package greet

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name")
	}

	message := fmt.Sprintf("Hi %v", name)

	return message, nil
}
