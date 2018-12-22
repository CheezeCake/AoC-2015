package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func stringIndices(s, substr string) []int {
	indices := []int{}
	start := 0
	for {
		i := strings.Index(s[start:], substr)
		if i < 0 {
			break
		}
		indices = append(indices, start+i)
		start += i + len(substr)
	}
	return indices
}

func possibleMolecules(molecule string, replacements map[string][]string) map[string]int {
	results := make(map[string]int)
	for from, tos := range replacements {
		for _, to := range tos {
			for _, index := range stringIndices(molecule, from) {
				result := molecule[0:index] + to + molecule[index+len(from):]
				results[result]++
			}
		}
	}
	return results
}

func makeMolecule(n int, current, target string, replacements map[string][]string, seen map[string]bool) bool {
	if _, ok := seen[current]; ok {
		return false
	}
	seen[current] = true
	if current == target {
		fmt.Println("part 2: ", n)
		return true
	}
	for result, _ := range possibleMolecules(current, replacements) {
		if makeMolecule(n+1, result, target, replacements, seen) {
			return true
		}
	}
	return false
}

func main() {
	replacements := make(map[string][]string)
	reverseReplacements := make(map[string][]string)
	medicine := ""

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "=>") {
			fields := strings.Split(line, " => ")
			from := fields[0]
			to := fields[1]
			if replacements[from] == nil {
				replacements[from] = []string{to}
			} else {
				replacements[from] = append(replacements[from], to)
			}
			if reverseReplacements[to] == nil {
				reverseReplacements[to] = []string{from}
			} else {
				reverseReplacements[to] = append(reverseReplacements[to], from)
			}

		} else if len(line) > 0 {
			medicine = line
		}
	}

	fmt.Println("part 1:", len(possibleMolecules(medicine, replacements)))
	makeMolecule(0, medicine, "e", reverseReplacements, make(map[string]bool))
}
