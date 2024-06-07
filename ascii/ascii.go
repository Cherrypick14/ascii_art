package ascii

import (
	"fmt"
)

// AsciiArt processes words, printing their ASCII art 
// character by character and adding new lines as needed.

func AsciiArt(words []string, contents2 []string) {
	countSpace := 0
	for _, word := range words {
		if word != "" {
			for i := 0; i < 8; i++ {
				for _, char := range word {
					if char == '\n' {
						continue
					}
					if !(char >= 32 && char <= 126) {
						fmt.Println("Error: Input contains non-ASCII characters")
						return
					}
					// Print the calculated index of 'char' Ascii Art in content2. 
					fmt.Print(contents2[int(char-' ')*9+1+i])
				}
				fmt.Println()
			}
		} else {
			countSpace++
			if countSpace < len(words) {
				fmt.Println()
			}
		}
	}
}
