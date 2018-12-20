package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const Time = 2503

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type reindeer struct {
	speed       int
	flyingTime  int
	restingTime int

	score    int
	distance int
}

func (r *reindeer) advance(currentTime int) {
	if currentTime%(r.flyingTime+r.restingTime) < r.flyingTime {
		r.distance += r.speed
	}
}

func simlulate(reindeers []*reindeer) (maxDistance, maxScore int) {
	for time := 0; time < Time; time++ {
		reindeersWithMaxDistance := []int{}

		for i, r := range reindeers {
			r.advance(time)
			if r.distance > maxDistance {
				maxDistance = r.distance
				reindeersWithMaxDistance = []int{i}
			} else if r.distance == maxDistance {
				reindeersWithMaxDistance = append(reindeersWithMaxDistance, i)
			}
		}

		for _, ri := range reindeersWithMaxDistance {
			reindeers[ri].score++
			maxScore = max(maxScore, reindeers[ri].score)
		}
	}
	return
}

func main() {
	reindeers := []*reindeer{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())

		speed, _ := strconv.Atoi(words[3])
		flyingTime, _ := strconv.Atoi(words[6])
		restingTime, _ := strconv.Atoi(words[len(words)-2])

		reindeers = append(reindeers,
			&reindeer{speed: speed, flyingTime: flyingTime, restingTime: restingTime})
	}

	maxDistance, maxScore := simlulate(reindeers)
	fmt.Println("part 1:", maxDistance)
	fmt.Println("part 2:", maxScore)
}
