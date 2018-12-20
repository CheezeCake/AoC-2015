package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func combinations(volume int, n int, containers []int, combinationCount map[int]int, minContainers *int) int {
	if volume == 0 {
		combinationCount[n]++
		if n < *minContainers {
			*minContainers = n
		}
		return 1
	}
	if volume < 0 {
		return 0
	}

	count := 0
	for i, c := range containers {
		count += combinations(volume-c, n+1, containers[i+1:], combinationCount, minContainers)
	}
	return count
}

func readContainers() []int {
	containers := []int{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		c, _ := strconv.Atoi(scanner.Text())
		containers = append(containers, c)
	}

	return containers
}

func main() {
	containers := readContainers()
	combinationCount := make(map[int]int)
	minContainers := math.MaxInt32

	fmt.Println("part 1:", combinations(150, 0, containers, combinationCount, &minContainers))
	fmt.Println("part 2:", combinationCount[minContainers])
}
