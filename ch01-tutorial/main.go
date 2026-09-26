package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command")
		return
	}

	switch os.Args[1] {
	case "1":
		helloWorld()
	default:
		fmt.Println("Unknown command")
	}

}

func helloWorld() {
	fmt.Println("Hello, World!")
}
