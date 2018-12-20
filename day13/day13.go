package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func computeChange(guest1, guest2 string, guests map[string]*guest) int {
	return guests[guest1].points[guest2] + guests[guest2].points[guest1]
}

func generate(currentSeat int, table []string, totalChange int, guests map[string]*guest, max *int) {
	if currentSeat == len(table) {
		totalChange += computeChange(table[0], table[len(table)-1], guests)
		if totalChange > *max {
			*max = totalChange
		}
		return
	}

	for name, guest := range guests {
		if guest.seated {
			continue
		}

		guest.seated = true
		table[currentSeat] = name

		change := totalChange
		if currentSeat > 0 {
			change += computeChange(table[currentSeat-1], name, guests)
		}
		generate(currentSeat+1, table, change, guests, max)

		guest.seated = false
	}
}

type guest struct {
	points map[string]int
	seated bool
}

func newGuest() *guest {
	return &guest{points: make(map[string]int)}
}

func solve(guests map[string]*guest) int {
	maxChange := 0
	table := make([]string, len(guests))
	generate(0, table, 0, guests, &maxChange)
	return maxChange
}

func main() {
	guests := make(map[string]*guest)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())

		subject := words[0]
		cod := words[len(words)-1]
		cod = cod[:len(cod)-1]
		points, _ := strconv.Atoi(words[3])

		if guests[subject] == nil {
			guests[subject] = newGuest()
		}
		if words[2] == "gain" {
			guests[subject].points[cod] = points
		} else {
			guests[subject].points[cod] = -points
		}
	}

	fmt.Println("part 1:", solve(guests))

	me := newGuest()
	for name, guest := range guests {
		guest.points["me"] = 0
		me.points[name] = 0
	}
	guests["me"] = me

	fmt.Println("part 2:", solve(guests))
}
