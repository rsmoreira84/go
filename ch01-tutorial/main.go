package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command")
		return
	}

	switch os.Args[1] {
	case "helloworld":
		helloWorld()
	case "echo1":
		showCommandArgs(os.Stdout)
	case "echo2":
		showCommandArgs2(os.Stdout)
	case "echo2b":
		showCommandArgs2b(os.Stdout)
	case "echo3":
		showCommandArgs3(os.Stdout)
	case "echo4":
		showCommandArgs(os.Stdout)
		fmt.Fprintln(os.Stdout, "----------")
		showCommandArgs2(os.Stdout)
		fmt.Fprintln(os.Stdout, "----------")
		showCommandArgs2b(os.Stdout)
		fmt.Fprintln(os.Stdout, "----------")
		showCommandArgs3(os.Stdout)
	default:
		fmt.Println("Unknown command")
	}

}

func helloWorld() {
	fmt.Println("Hello, World!")
}

// showCommandArgs prints all command-line arguments passed to the program.
func showCommandArgs(w io.Writer) {
	start := time.Now()
	fmt.Fprintln(w, "ECHO 1:")
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Fprintln(w, s)
	fmt.Fprintln(w, "Execution time:", time.Since(start))
}

// showCommandArgs2 prints all command-line arguments passed to the program using a range loop.
func showCommandArgs2(w io.Writer) {
	start := time.Now()
	fmt.Fprintln(w, "ECHO 2:")
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintln(w, s)
	fmt.Fprintln(w, "Execution time:", time.Since(start))
}

// showCommandArgs2b prints all command-line arguments passed to the program using a range loop.
// getting also the index from the range command.
// Range always provides both the index and the value, even if the index is ignored with an underscore.
func showCommandArgs2b(w io.Writer) {
	start := time.Now()
	fmt.Fprintln(w, "ECHO 2B:")
	s, sep := "", ""

	for i, arg := range os.Args[1:] {
		s += sep + fmt.Sprintf("[%d] %s", i, arg)
		sep = "\n"
	}
	fmt.Fprintln(w, s)
	fmt.Fprintln(w, "Execution time:", time.Since(start))
}

func showCommandArgs3(w io.Writer) {
	start := time.Now()
	fmt.Fprintln(w, "ECHO 3:")
	fmt.Fprintln(w, strings.Join(os.Args[1:], " "))
	fmt.Fprintln(w, "Execution time:", time.Since(start))
}
