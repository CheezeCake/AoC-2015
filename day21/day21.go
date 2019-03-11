package main

import (
	"fmt"
	"math"
)

const (
	Player = 0
	Boss   = 1
)

type character struct {
	hp, damage, armor int
}

type item struct {
	name                string
	cost, damage, armor int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func winner(player, boss character) int {
	playerActualDamage := max(player.damage-boss.armor, 1)
	bossActualDamage := max(boss.damage-player.armor, 1)

	playerRounds := player.hp / bossActualDamage
	playerHpLeft := player.hp % bossActualDamage
	if playerHpLeft > 0 {
		playerRounds++
	}
	bossRounds := boss.hp / playerActualDamage
	bossHpLeft := boss.hp % playerActualDamage
	if bossHpLeft > 0 {
		bossRounds++
	}

	if playerRounds > bossRounds {
		return Player
	}
	if bossRounds > playerRounds {
		return Boss
	}
	if bossHpLeft == 0 {
		return Player
	}
	if playerHpLeft == 0 {
		return Boss
	}
	return Player
}

// func winner(player, boss character) int {
// 	playerActualDamage := max(player.damage-boss.armor, 1)
// 	bossActualDamage := max(boss.damage-player.armor, 1)

// 	for {
// 		boss.hp -= playerActualDamage
// 		if boss.hp <= 0 {
// 			return Player
// 		}

// 		player.hp -= bossActualDamage
// 		if player.hp <= 0 {
// 			return Boss
// 		}
// 	}
// }

func fight(weapon, armor, ring1, ring2 item) bool {
	player := character{
		100,
		weapon.damage + armor.damage + ring1.damage + ring2.damage,
		weapon.armor + armor.armor + ring1.armor + ring2.armor,
	}
	boss := character{104, 8, 1}

	return (winner(player, boss) == Player)
}

func main() {
	weapons := []item{
		{"Dagger", 8, 4, 0},
		{"Shortsword", 10, 5, 0},
		{"Warhammer", 25, 6, 0},
		{"Longsword", 40, 7, 0},
		{"Greataxe", 74, 8, 0},
	}
	armors := []item{
		{"No Armor", 0, 0, 0},
		{"Leather", 13, 0, 1},
		{"Chainmail", 31, 0, 2},
		{"Splintmail", 53, 0, 3},
		{"Bandedmail", 75, 0, 4},
		{"Platemail", 102, 0, 5},
	}
	rings := []item{
		{"No Ring", 0, 0, 0},
		{"Damage +1 ", 25, 1, 0},
		{"Damage +2 ", 50, 2, 0},
		{"Damage +3 ", 100, 3, 0},
		{"Defense +1", 20, 0, 1},
		{"Defense +2", 40, 0, 2},
		{"Defense +3", 80, 0, 3},
	}

	minCostToWin := math.MaxInt32
	maxCostToLose := 0

	for _, weapon := range weapons {
		for _, armor := range armors {
			// no ring
			totalCost := weapon.cost + armor.cost
			if fight(weapon, armor, item{}, item{}) {
				minCostToWin = min(minCostToWin, totalCost)
			} else {
				maxCostToLose = max(maxCostToLose, totalCost)
			}

			for _, ring1 := range rings {
				for _, ring2 := range rings[1:] {
					if ring1.name == ring2.name {
						continue
					}

					totalCost := weapon.cost + armor.cost + ring1.cost + ring2.cost
					if fight(weapon, armor, ring1, ring2) {
						minCostToWin = min(minCostToWin, totalCost)
					} else {
						maxCostToLose = max(maxCostToLose, totalCost)
					}
				}
			}
		}
	}

	fmt.Println("part 1:", minCostToWin)
	fmt.Println("part 2:", maxCostToLose)
}
