package common

import (
	"bufio"
	"log"
	"os"
)

// ReadStdin reads content from stdin into a string
func ReadStdin() string {
	inval := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inval += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return inval
}

// ReadStdinLines reads lines from stdin into a slice of strings
func ReadStdinLines() []string {
	var inval []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inval = append(inval, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return inval
}
