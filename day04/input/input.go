package main

// Multiple import.
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Read line by using scanner.
	// Initialize scanner listen to standard in.
	scanner := bufio.NewScanner(os.Stdin)
	// Waiting for next token (line).
	scanner.Scan()
	// Get string from the scanner.
	// This will remove `\n` or '\r\n` for you.
	fmt.Println(scanner.Text())

	// Read line by using reader.
	reader := bufio.NewReader(os.Stdin)
	// Read string until the end at `\n`.
	// Will return (string text, error message)
	text, err := reader.ReadString('\n')
	// Check if there is an error.
	if err != nil {
		fmt.Println(err)
	} else {
		// You should remove `\n` or `\r\n`(in windows)
		// by yourself.
		text = strings.Replace(text, "\n", "", 1)
		fmt.Println(text)
	}
}
