package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func newPoint(data string) point {
	coords := strings.Split(data, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	return point{x: x, y: y}
}

type action struct {
	action   string
	from, to point
}

func parseAction(line string) action {
	words := strings.Split(line, " ")
	var act, fromStr string
	toStr := words[len(words)-1]

	if words[0] == "turn" {
		act = words[1]
		fromStr = words[2]
	} else {
		act = words[0]
		fromStr = words[1]
	}

	return action{
		action: act,
		from:   newPoint(fromStr),
		to:     newPoint(toStr),
	}
}

type Grid struct {
	grid [1000][1000]struct {
		on         bool
		brightness int
	}

	lightsOn        int
	totalBrightness int
}

func (g *Grid) do(act action) {
	for y := act.from.y; y <= act.to.y; y++ {
		for x := act.from.x; x <= act.to.x; x++ {
			switch act.action {
			case "on":
				if !g.grid[y][x].on {
					g.lightsOn++
					g.grid[y][x].on = true
				}
				g.grid[y][x].brightness++
				g.totalBrightness++
			case "off":
				if g.grid[y][x].on {
					g.lightsOn--
					g.grid[y][x].on = false
				}
				if g.grid[y][x].brightness > 0 {
					g.grid[y][x].brightness--
					g.totalBrightness--
				}
			case "toggle":
				if g.grid[y][x].on {
					g.lightsOn--
					g.grid[y][x].on = false
				} else {
					g.lightsOn++
					g.grid[y][x].on = true
				}
				g.grid[y][x].brightness += 2
				g.totalBrightness += 2
			}
		}
	}
}

func main() {
	var grid Grid

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		grid.do(parseAction(scanner.Text()))
	}

	fmt.Println("part 1:", grid.lightsOn)
	fmt.Println("part 2:", grid.totalBrightness)
}
