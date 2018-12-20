package main

import "fmt"

func main() {
	var input string
	fmt.Scanln(&input)

	floor := 0
	pos := -1
	for i, c := range input {
		if c == '(' {
			floor++
		} else {
			floor--
		}

		if floor == -1 && pos == -1 {
			pos = i + 1
		}
	}

	fmt.Println("part 1:", floor)
	fmt.Println("part 2:", pos)
}
