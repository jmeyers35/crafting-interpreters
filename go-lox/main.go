package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jmeyers35/golox/pkg/scanner"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("usage: golox [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		RunFile(os.Args[1])
	} else {
		RunPrompt()
	}
}

func RunPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		bytes, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Printf("error reading input: %v\n", err)
		}

		run(bytes)
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
	s := scanner.New()
	tokens := s.Scan(bytes)

	for _, token := range tokens {
		fmt.Println(token)
	}

	return nil
}
