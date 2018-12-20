package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	codeChars := 0
	valueChars := 0
	encodedChars := 0
	for scanner.Scan() {
		str := scanner.Text()
		codeChars += len(str)
		encodedChars += len(str)

		for i := 0; i < len(str); i++ {
			if str[i] == '\\' {
				encodedChars++
				if str[i+1] == '\\' || str[i+1] == '"' {
					encodedChars++
					i++
				} else { // hex
					i += 3
				}
			}
			valueChars++
		}

		// don't count the two surrounding "
		valueChars -= 2
		// two \ to encode the original surrounding " + two new surrounding "
		encodedChars += 4
	}

	fmt.Println("part 1:", codeChars-valueChars)
	fmt.Println("part 2:", encodedChars-codeChars)
}
