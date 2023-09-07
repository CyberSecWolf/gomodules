package drake

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input(str string) (string, error) { // requirements: bufio, fmt, os, strings
	var user_input string

	reader := bufio.NewReader(os.Stdin) // Creates a new reader to read from standard input

	fmt.Print(str) // Prompts the user for input

	user_input, err := reader.ReadString('\n') // Read the input from the user
	if err != nil {
		return "Error: ", err
	}

	return strings.TrimSpace(user_input), err
}
