package main

import (
	"fmt"
	"os"
	"strings"

	"ascii/ascii"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: Incorrect number of arguments ", len(os.Args))
		return
	}
	// Get an input text from the commandline.
	inputtext := os.Args[1]

	// Check if there are any args in the commandline and if the input is also empty.
	if inputtext == "" {
		return
	}
	// Check if inputtext is a newline character, if it is, then print a newline.
	if inputtext == "\\n" {
		fmt.Print("\n")
		return
	}
	if inputtext == "\\a" || inputtext == "\\0" || inputtext == "\\f" || inputtext == "\\v" || inputtext == "\\r" {
		fmt.Println("Error: Non printable character ", inputtext)
		return
	}
	inputtext = strings.ReplaceAll(inputtext, "\\t", "    ")
	inputtext = strings.ReplaceAll(inputtext, "\\b", "\b")

	// Logic process for handlng the backspace.
	for i := 0; i < len(inputtext); i++ {
		indexb := strings.Index(inputtext, "\b")

		if indexb > 0 {
			inputtext = inputtext[:indexb-1] + inputtext[indexb+1:]
		}
	}

	// Split our input text to a string slice and separate with a newline.
	words := strings.Split(inputtext, "\\n")

	// Read content from one of the banner files.
	contents, err := os.ReadFile("standard.txt")
	// Apply error handling if file is incorrect.
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	fileInfo, err := os.Stat("standard.txt")
	if err != nil {
		fmt.Println("Error reading file information", err)
		return
	}
	fileSize := fileInfo.Size()

	if fileSize != 6623 {
		fmt.Println("Error with the file size", fileSize)
		return
	}

	// Split the content to a string slice and separate with newline.
	contents2 := strings.Split(string(contents), "\n")

	ascii.AsciiArt(words, contents2)
}
