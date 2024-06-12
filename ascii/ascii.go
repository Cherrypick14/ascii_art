package ascii

import (
	// "fmt"
	"strings"
)

// AsciiArt processes words, printing their ASCII art
// character by character and adding new lines as needed.

func AsciiArt(words []string, contents2 []string) string {

	var result strings.Builder

	countSpace := 0
	for _, word := range words {
		if word != "" {
			for i := 0; i < 8; i++ {
				for _, char := range word {
					if char == '\n' {
						continue
					}
					if !(char >= 32 && char <= 126) {
						result.WriteString("Error: Input contains non-ASCII characters")	
						
					}
					// Print the calculated index of 'char' Ascii Art in content2. 
					result.WriteString(contents2[int(char-' ')*9+1+i])
				}
				result.WriteString("\n")
			}
		} else {
			countSpace++
			if countSpace < len(words) {
				result.WriteString("\n")
			}
		}
	}
	return result.String()
}
