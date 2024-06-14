package ascii

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestAsciiArt(t *testing.T) {
	type parameters struct {
		words     []string
		contents2 []string
	}
	type testCases struct {
		name     string
		args     parameters
		expected string
	}
	// We read the bannerfile directly to save on space/make the code readable
	// otherwise we would have provided the contents in banner file as a []string.
	contents, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Print("Error reading from file", err)
	}
	new_Expected, err := os.ReadFile("expected.txt")
	if err != nil {
		fmt.Println("Error reading from file", err)
	}

	textArt := strings.Split(string(contents), "\n")

	tests := []testCases{
		{name: "test empty", args: parameters{words: []string{}, contents2: textArt}, expected: ""},
		{name: "test newline(\n)", args: parameters{words: []string{""}, contents2: textArt}, expected: ""},
		{name: "test Non ascii", args: parameters{words: []string{"你好"}, contents2: textArt}, expected: "Error: Input contains non-ASCII characters"},
		{name: "test hello", args: parameters{words: []string{"hello"}, contents2: textArt}, expected: string(new_Expected)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := AsciiArt(tt.args.words, tt.args.contents2)
			if actual != tt.expected {
				t.Errorf("Expected different output, got: %s", actual)
			}

		})
	}
}
