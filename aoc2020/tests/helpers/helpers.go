package helpers

import (
	"io/ioutil"
	"log"
	"strings"
)

// LoadFixture loads a file into a slice of strings,
// one string per line
func LoadFixture(filepath string) []string {
	data := LoadFixtureToString(filepath)
	lines := strings.Split(data, "\n")
	return lines
}

// LoadFixtureToString loads a file into a single string
func LoadFixtureToString(filepath string) string {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
