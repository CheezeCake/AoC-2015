package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func paperNeeded(l, w, h int) int {
	face1 := l * w
	face2 := w * h
	face3 := h * l
	minFace := min(min(face1, face2), face3)

	return 2*(face1+face2+face3) + minFace
}

func ribbonNeeded(l, w, h int) int {
	face1 := l + w
	face2 := w + h
	face3 := h + l
	minFace := min(min(face1, face2), face3)

	return (2 * minFace) + (l * w * h)
}

func main() {
	paper := 0
	ribbon := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Split(line, "x")
		l, _ := strconv.Atoi(fields[0])
		w, _ := strconv.Atoi(fields[1])
		h, _ := strconv.Atoi(fields[2])

		paper += paperNeeded(l, w, h)
		ribbon += ribbonNeeded(l, w, h)
	}

	fmt.Println("part 1:", paper)
	fmt.Println("part 2:", ribbon)
}
