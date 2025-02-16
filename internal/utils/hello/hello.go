package hello

import "fmt"

var unusedVar = 42 // Unused variable

func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	if false {
		fmt.Println("This will be printed")
	}
	return "Hello, " + name + "!"
}
