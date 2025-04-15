package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("usage: golox [script]")
		os.Exit(64)
	} else if len(os.Args) == 1 {
		RunFile(os.Args[0])
	} else {
		RunPrompt()
	}
}

func RunPrompt() {
	for {
		fmt.Print("> ")
		buf := make([]byte, 0, 1024)
		_, err := os.Stdin.Read(buf)
		if err != nil {
			fmt.Printf("error reading input: %v\n", err)
		}

		run(buf)
	}
}

func RunFile(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}

	return run(bytes)
}

func run(bytes []byte) error {
	panic("unimplmented")
}
