package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func deliver(house, factor, maxDeliveries int, deliveries *map[int]int) int {
	if (*deliveries)[house] == maxDeliveries {
		return 0
	}
	(*deliveries)[house]++
	return (house * factor)
}

func deliverPresents(input, factor, maxDeliveries int) int {
	deliveries := make(map[int]int)

	for house := 2; ; house++ {
		presents := 0

		for elf := 1; elf*elf <= house; elf++ {
			if house%elf == 0 {
				presents += deliver(elf, factor, maxDeliveries, &deliveries)
				quotient := house / elf
				if quotient != elf {
					presents += deliver(quotient, factor, maxDeliveries, &deliveries)
				}
			}
		}

		if presents >= input {
			return house
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage:", os.Args[0], "input")
		os.Exit(1)
	}

	input, _ := strconv.Atoi(os.Args[1])

	fmt.Println("part 1:", deliverPresents(input, 10, math.MaxInt32))
	fmt.Println("part 2:", deliverPresents(input, 11, 50))
}
