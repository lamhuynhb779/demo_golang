package package_module1

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v", name)
	return message
}