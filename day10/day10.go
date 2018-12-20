package main

import "fmt"

const input = "1113122113"

func lookAndSay(str string) string {
	res := ""
	i := 0
	for i < len(str) {
		j := i
		for ; j < len(str) && str[i] == str[j]; j++ {
		}
		n := j - i
		res += string('0'+n) + string(str[i])
		i = j
	}
	return res
}

func main() {
	str := input
	for i := 0; i < 40; i++ {
		str = lookAndSay(str)
	}
	fmt.Println("part 1:", len(str))
	for i := 0; i < 10; i++ {
		str = lookAndSay(str)
	}
	fmt.Println("part 2:", len(str))
}
