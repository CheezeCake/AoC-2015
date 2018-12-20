package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

const input = "bgvyzdsv"

func main() {
	part1Done := false
	part2Done := false
	for i := 1; !part1Done || !part2Done; i++ {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(input+strconv.Itoa(i))))
		if !part1Done && strings.HasPrefix(hash, "00000") {
			fmt.Println("part 1:", i)
			part1Done = true
		}
		if !part2Done && strings.HasPrefix(hash, "000000") {
			fmt.Println("part 2:", i)
			part2Done = true
		}
	}
}
