package main

import (
	"fmt"
	"os"
)

func main() {
	// Check for command-line arguments.
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: go run main.go [major|minor|patch]")
		os.Exit(1)
	}

	arg := os.Args[1]

	// Ensure the argument provided is either 'major', 'minor', or 'patch'.
	switch arg {
	case "major", "minor", "patch":
		fmt.Println(arg) // Print to standard output
	default:
		fmt.Fprintf(os.Stderr, "invalid argument: %s\n", arg) // Print error message to standard error output
		fmt.Fprintln(os.Stderr, "Please specify one of the following: major, minor, or patch.")
		os.Exit(2) // Exit the program with a different exit code to indicate invalid argument was provided
	}
}
