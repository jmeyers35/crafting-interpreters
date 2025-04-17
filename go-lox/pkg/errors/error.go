package errors

import (
	"log"
	"os"
)

func LogError(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	ll := log.New(os.Stderr, "", 0)
	ll.Printf("[line %d] Error %s: %s", line, where, message)
}
