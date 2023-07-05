package main

import (
	"fmt"
	"hayashi/module1" // import module
	"hayashi/module3/package_module3" // import package (syntax <module-name>/<package-name>)
	"hayashi/module3/package2_module3" // import package
)

func main() {
	message := package_module1.Hello("Lam")

	fmt.Println(message)

	fmt.Println(package_module3.Goodbye("Lam"))

	fmt.Println(package2_module3.Say("Lam"))
}
