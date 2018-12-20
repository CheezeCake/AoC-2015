package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
)

func _nice1(str string) bool {
	vowels := 0
	twice := 0
	invalidStrs := false

	for i, c := range str {
		for _, v := range "aeiou" {
			if c == v {
				vowels++
				break
			}
		}

		if i > 0 {
			if str[i-1] == byte(c) {
				twice++
			}
			if str[i-1] == 'a' && c == 'b' {
				invalidStrs = true
			} else if str[i-1] == 'c' && c == 'd' {
				invalidStrs = true
			} else if str[i-1] == 'p' && c == 'q' {
				invalidStrs = true
			} else if str[i-1] == 'x' && c == 'y' {
				invalidStrs = true
			}
		}
		if invalidStrs {
			break
		}
	}

	return (!invalidStrs && vowels >= 3 && twice > 0)
}

func matches(pattern, str string) bool {
	re, _ := pcre.Compile(pattern, 0)
	return re.MatcherString(str, 0).Matches()
}

func nice1(str string) bool {
	if !matches(`[aeiou].*[aeiou].*[aeiou]`, str) {
		return false
	}
	if !matches(`(.)\1`, str) {
		return false
	}
	for _, substr := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(str, substr) {
			return false
		}
	}
	return true
}

func nice2(str string) bool {
	return matches(`(..).*\1`, str) && matches(`(.).\1`, str)
}

func main() {
	n1 := 0
	n2 := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		if nice1(str) {
			n1++
		}
		if nice2(str) {
			n2++
		}

	}
	fmt.Println("part 1:", n1)
	fmt.Println("part 2:", n2)
}
