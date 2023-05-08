package builtins

import "fmt"

// create a static variable to hold the history
var history []string

func SaveHistory(args ...string ) error {
	history = append(history, args...)
	return nil
}

func PrintHistory() error {
	for _, arg := range history {
		fmt.Println(arg)
	}
	return nil
}
