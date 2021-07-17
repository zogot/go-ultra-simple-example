package greeter

import "fmt"

// Greet can be accessed outside the packager 'greeter'
func Greet() {
	fmt.Println("Hello World!")
}
