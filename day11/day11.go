package main

import (
	"fmt"
	"strings"
)

const (
	Input          = "vzbxkghx"
	ForbiddenChars = "iol"
)

func hasIncreasingStraight(pwd string) bool {
	for i := 0; i < len(pwd)-2; i++ {
		if pwd[i+1] == pwd[i]+1 && pwd[i+2] == pwd[i]+2 {
			return true
		}
	}
	return false
}

func hasPairs(pwd string) bool {
	pairChar := make(map[byte]bool)
	for i := 0; i < len(pwd)-1; i++ {
		if pwd[i] == pwd[i+1] {
			pairChar[pwd[i]] = true
			i++
		}
	}
	return len(pairChar) >= 2
}

func hasForbiddenChars(pwd string) bool {
	return strings.IndexAny(pwd, ForbiddenChars) >= 0
}

func valid(pwd string) bool {
	return hasIncreasingStraight(pwd) && hasPairs(pwd) && !hasForbiddenChars(pwd)
}

func increment(pwd string) string {
	res := []byte(pwd)
	if pos := strings.IndexAny(pwd, ForbiddenChars); pos >= 0 {
		res[pos] = 'a' + ((res[pos] - 'a' + 1) % 26)
		for i := pos + 1; i < len(pwd); i++ {
			res[i] = 'a'
		}
	} else {
		carry := 1
		for i := len(pwd) - 1; i >= 0 && carry > 0; i-- {
			l := pwd[i] - 'a' + byte(carry)
			if l >= 26 {
				carry = 1
				l %= 26
			} else {
				carry = 0
			}
			res[i] = 'a' + l
		}
	}
	return string(res)
}

func next(pwd string) string {
	pwd = increment(pwd)
	for !valid(pwd) {
		pwd = increment(pwd)
	}
	return pwd
}

func main() {
	pwd := next(Input)
	fmt.Println("part 1:", pwd)
	pwd = next(pwd)
	fmt.Println("part 2:", pwd)
}
