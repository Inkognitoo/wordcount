package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := readInput()
	count := countWords(str)

	fmt.Println(count)
}

// readInput reads command line arguments and returns them.
func readInput() string {
	if len(os.Args) < 2 {
		return ""
	}

	return os.Args[1]
}

// countWords gets string and return numbers of words in it
func countWords(words string) int {
	sliceWords := strings.Fields(words)

	return len(sliceWords)
}
