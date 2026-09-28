package main

import (
	"bufio"
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

	// Hello World exercise
	case "helloworld":
		helloWorld()

	// Show command-line arguments exercises
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

	// Duplicated lines exercises
	case "dup1":
		duplicatedLines1(os.Stdout)

	case "dup2":
		duplicatedLines2(os.Stdout)

	case "dup3":
		duplicatedLines3(os.Stdout)

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

// duplicatedLines1 prints all lines from the standard input that appear more than once, along with their counts.
// This function reads lines from the standard input and counts their occurrences.
// It then prints only the lines that appear more than once, along with their counts.
// So, the user will need to type on the terminal, and to send the input, they can use Ctrl+D (EOF) to signal the end of input.
func duplicatedLines1(w io.Writer) {
	start := time.Now()

	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	fmt.Fprintln(w, "Execution time:", time.Since(start))
}

// duplicatedLines2 prints all lines from the specified files (or standard input if no files are specified) that appear more than once, along with their counts.
// This function reads lines from the specified files (or standard input) and counts their occurrences.
// It then prints only the lines that appear more than once, along with their counts.
// To run it the user needs to inform the file to be checked after the command name, for example:
//
//	go run main.go dup2 duplicated_lines_file.md
func duplicatedLines2(w io.Writer) {
	start := time.Now()

	counts := make(map[string]int)
	files := os.Args[2:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Fprintf(w, "%d\t%s\n", n, line)
		}
	}

	fmt.Fprintln(w, "Execution time: ", time.Since(start))
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

func duplicatedLines3(w io.Writer) {
	start := time.Now()

	fmt.Fprintln(w, "Execution time: ", time.Since(start))
}
