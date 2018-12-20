package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

const Size = 100

type grid [Size][Size]byte

func (g *grid) lightsOn() int {
	cnt := 0
	for _, row := range g {
		cnt += bytes.Count(row[:], []byte{'#'})
	}
	return cnt
}

func (g *grid) neighborsOn(x, y int) int {
	directions := [8]struct{ x, y int }{
		{-1, 0},  // left
		{-1, -1}, // top-left
		{0, -1},  // top
		{1, -1},  // top-right
		{1, 0},   // right
		{1, 1},   // bottom-right
		{0, 1},   // bottom
		{-1, 1},  // bottom-left
	}

	cnt := 0
	for _, dir := range directions {
		nx := x + dir.x
		ny := y + dir.y
		if nx >= 0 && nx < Size && ny >= 0 && ny < Size {
			if g[ny][nx] == '#' {
				cnt++
			}
		}
	}
	return cnt
}

func (g *grid) step() {
	newGrid := grid{}
	for y := 0; y < Size; y++ {
		for x := 0; x < Size; x++ {
			no := g.neighborsOn(x, y)
			if g[y][x] == '#' && no != 2 && no != 3 {
				newGrid[y][x] = '.'
			} else if g[y][x] == '.' && no == 3 {
				newGrid[y][x] = '#'
			} else {
				newGrid[y][x] = g[y][x]
			}
		}
	}
	*g = newGrid
}

func (g *grid) turnOnCorners() {
	g[0][0] = '#'
	g[0][Size-1] = '#'
	g[Size-1][0] = '#'
	g[Size-1][Size-1] = '#'
}

func (g *grid) print() {
	for _, row := range g {
		fmt.Printf("%s\n", row)
	}
}

func (g *grid) read(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for i := 0; scanner.Scan() && i < Size; i++ {
		copy(g[i][:], scanner.Text()[:])
	}
}

func main() {
	grid := grid{}
	grid.read(os.Stdin)

	g1 := grid
	for i := 1; i <= 100; i++ {
		g1.step()
	}
	fmt.Println("part 1:", g1.lightsOn())

	grid.turnOnCorners()
	for i := 1; i <= 100; i++ {
		grid.step()
		grid.turnOnCorners()
	}
	fmt.Println("part 2:", grid.lightsOn())
}
