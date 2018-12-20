package main

import "fmt"

var directions = map[byte][2]int{
	'<': {-1, 0},
	'^': {0, -1},
	'>': {1, 0},
	'v': {0, 1},
}

type house struct {
	x, y int
}

func (h house) next(direction byte) house {
	return house{
		x: h.x + directions[direction][0],
		y: h.y + directions[direction][1],
	}
}

func solvePart1(input string) int {
	visited := make(map[house]bool)
	currentHouse := house{0, 0}

	visited[currentHouse] = true
	for _, d := range input {
		currentHouse = currentHouse.next(byte(d))
		visited[currentHouse] = true
	}

	return len(visited)
}

func solvePart2(input string) int {
	visited := make(map[house]bool)
	santaCurrentHouse := house{0, 0}
	robotCurrentHouse := house{0, 0}

	visited[santaCurrentHouse] = true
	for i := 0; i < len(input)-1; i += 2 {
		santaCurrentHouse = santaCurrentHouse.next(input[i])
		visited[santaCurrentHouse] = true
		robotCurrentHouse = robotCurrentHouse.next(input[i+1])
		visited[robotCurrentHouse] = true
	}

	return len(visited)
}

func main() {
	var input string
	fmt.Scanln(&input)

	fmt.Println("part 1:", solvePart1(input))
	fmt.Println("part 2:", solvePart2(input))
}
