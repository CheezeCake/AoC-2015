package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ingredient struct {
	name       string
	properties map[string]int
}

func parsePropertyValue(field string) int {
	value, _ := strconv.Atoi(strings.TrimSuffix(field, ","))
	return value
}

func parseIngredient(line string) ingredient {
	i := ingredient{}
	words := strings.Fields(line)

	i.name = strings.TrimSuffix(words[0], ":")
	i.properties = make(map[string]int)
	i.properties["capacity"] = parsePropertyValue(words[2])
	i.properties["durability"] = parsePropertyValue(words[4])
	i.properties["flavor"] = parsePropertyValue(words[6])
	i.properties["texture"] = parsePropertyValue(words[8])
	i.properties["calories"] = parsePropertyValue(words[10])

	return i
}

func solve(currentIngredient int, ammountLeft int, propertiesSum map[string]int, ingredients []ingredient, maxScore1 *int, maxScore2 *int) {
	if currentIngredient == len(ingredients) {
		score := 1
		for property, sum := range propertiesSum {
			if property != "calories" {
				if sum < 0 {
					return
				}
				score *= sum
			}
		}
		if score > *maxScore1 {
			*maxScore1 = score
		}
		if propertiesSum["calories"] == 500 && score > *maxScore2 {
			*maxScore2 = score
		}
		return
	}

	for ammount := 0; ammount <= ammountLeft; ammount++ {
		for property, value := range ingredients[currentIngredient].properties {
			propertiesSum[property] += value * ammount
		}
		solve(currentIngredient+1, ammountLeft-ammount, propertiesSum, ingredients,
			maxScore1, maxScore2)
		for property, value := range ingredients[currentIngredient].properties {
			propertiesSum[property] -= value * ammount
		}
	}
}

func main() {
	ingredients := []ingredient{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ingredients = append(ingredients, parseIngredient(scanner.Text()))
	}

	maxScore1 := 0
	maxScore2 := 0
	solve(0, 100, make(map[string]int), ingredients, &maxScore1, &maxScore2)

	fmt.Println("part 1:", maxScore1)
	fmt.Println("part 2:", maxScore2)
}
