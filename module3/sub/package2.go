package package2_module3

import "fmt"

func Say(name string) string {
	var message string
	message = fmt.Sprintf("Nice to meet you, %v", name)
	return message
}
