package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func generateRoutes(from string, cost int, distances *map[string]map[string]int, visited map[string]bool, min *int, max *int) {
	if len(visited) == len(*distances) {
		if cost < *min {
			*min = cost
		}
		if cost > *max {
			*max = cost
		}
		return
	}

	for to, d := range (*distances)[from] {
		if _, ok := visited[to]; ok {
			continue
		}

		visited[to] = true
		generateRoutes(to, cost+d, distances, visited, min, max)
		delete(visited, to)
	}
}

func minMaxRoute(distances *map[string]map[string]int) (min, max int) {
	min = math.MaxInt32
	max = 0
	for start, _ := range *distances {
		generateRoutes(start, 0, distances, map[string]bool{start: true}, &min, &max)
	}
	return
}

func main() {
	distances := make(map[string]map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		from := words[0]
		to := words[2]
		distance, _ := strconv.Atoi(words[4])

		if distances[from] == nil {
			distances[from] = make(map[string]int)
		}
		distances[from][to] = distance

		if distances[to] == nil {
			distances[to] = make(map[string]int)
		}
		distances[to][from] = distance
	}

	min, max := minMaxRoute(&distances)
	fmt.Println("part 1:", min)
	fmt.Println("part 2:", max)
}
