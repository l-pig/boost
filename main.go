package main

import (
	"boost/cmd"
	"fmt"
)

func main() {
	var z = 1
	fn := func() {
		z = z - 1
	}
	fn()
	b := 100 / z
	fmt.Print("Enter choice: ", b)

	cmd.Execute()
}
