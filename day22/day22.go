package main

import (
	"fmt"
	"math"
)

type spellCallback func(*Player, *Boss)

type spell struct {
	name   string
	mana   int
	f      spellCallback
	effect effect
}

type effect struct {
	timer          int
	tickCb, stopCb spellCallback
}

var spells = []spell{
	{
		"MagicMissile",
		53,
		func(p *Player, b *Boss) { b.hp -= 4 },
		effect{},
	},
	{
		"Drain",
		73,
		func(p *Player, b *Boss) { b.hp -= 2; p.hp += 2 },
		effect{},
	},
	{
		"Shield",
		113,
		func(p *Player, b *Boss) { p.armor += 7 },
		effect{
			timer:  6,
			stopCb: func(p *Player, b *Boss) { p.armor -= 7 },
		},
	},
	{
		"Poison",
		173,
		nil,
		effect{timer: 6, tickCb: func(p *Player, b *Boss) { b.hp -= 3 }},
	},
	{
		"Recharge",
		229,
		nil,
		effect{timer: 5, tickCb: func(p *Player, b *Boss) { p.mana += 101 }},
	},
}

type Player struct {
	hp, armor, mana int
	activeEffects   map[string]effect
}

func (p *Player) copy() Player {
	newPlayer := *p

	newPlayer.activeEffects = make(map[string]effect)
	for name, e := range p.activeEffects {
		newPlayer.activeEffects[name] = e
	}

	return newPlayer
}

func (p *Player) spellEffectIsActive(s spell) bool {
	_, ok := p.activeEffects[s.name]
	return ok
}

func (p *Player) castSpell(s spell, b *Boss) bool {
	if p.mana < s.mana || p.spellEffectIsActive(s) {
		return false
	}

	p.mana -= s.mana
	if s.f != nil {
		s.f(p, b)
	}
	if s.effect.timer > 0 {
		p.activeEffects[s.name] = s.effect
	}

	return true
}

func (p *Player) effectsTick(b *Boss) {
	for name, e := range p.activeEffects {
		if e.tickCb != nil {
			e.tickCb(p, b)
		}
		e.timer--
		if e.timer == 0 {
			if e.stopCb != nil {
				e.stopCb(p, b)
			}
			delete(p.activeEffects, name)
		} else {
			p.activeEffects[name] = e
		}
	}
}

type Boss struct {
	hp, damage int
}

func (b Boss) attack(p *Player) {
	p.hp -= max(b.damage-p.armor, 1)
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

func checkBossDead(b Boss, manaSpent int, minManaToWin *int) bool {
	if b.hp <= 0 {
		*minManaToWin = min(*minManaToWin, manaSpent)
		return true
	}
	return false
}

func turn(p_ Player, b_ Boss, manaSpent int, minManaToWin *int, hard bool) {
	if hard {
		p_.hp--
		if p_.hp <= 0 {
			return
		}
	}
	p_.effectsTick(&b_)
	if checkBossDead(b_, manaSpent, minManaToWin) {
		return
	}

	for _, spell := range spells {
		player := p_.copy()
		boss := b_

		if player.castSpell(spell, &boss) {
			mana := manaSpent + spell.mana
			if checkBossDead(boss, mana, minManaToWin) {
				continue
			}

			player.effectsTick(&boss)
			if checkBossDead(boss, mana, minManaToWin) {
				continue
			}
			boss.attack(&player)
			if player.hp <= 0 {
				continue
			}

			if mana < *minManaToWin {
				turn(player, boss, mana, minManaToWin, hard)
			}
		}
	}
}

func main() {
	player := Player{50, 0, 500, map[string]effect{}}
	boss := Boss{55, 8}

	minManaToWin := math.MaxInt32
	turn(player, boss, 0, &minManaToWin, false)
	fmt.Println("part 1:", minManaToWin)

	minManaToWin = math.MaxInt32
	turn(player, boss, 0, &minManaToWin, true)
	fmt.Println("part 2:", minManaToWin)
}
