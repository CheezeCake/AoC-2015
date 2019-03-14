package main

import (
	"fmt"
	"math"
)

type group struct {
	pkgs   []int
	weight int
}

func newGroup() group {
	return group{[]int{}, 0}
}

func (g *group) QE() uint64 {
	p := uint64(1)
	for _, v := range g.pkgs {
		p *= uint64(v)
	}
	return p
}

func (g *group) addPkg(pkg int) group {
	return group{append(g.pkgs, pkg), g.weight + pkg}
}

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func makeOtherGroups(curPkg int, packages, weights []int, targetWeight int) bool {
	if curPkg == len(packages) {
		for _, w := range weights {
			if w != targetWeight {
				return false
			}
		}
		return true
	}

	for i := range weights {
		if weights[i]+packages[curPkg] <= targetWeight {
			weights[i] += packages[curPkg]
			if makeOtherGroups(curPkg+1, packages, weights, targetWeight) {
				return true
			}
			weights[i] -= packages[curPkg]
		}
	}

	return false
}

func solve(curPkg int, packages []int, firstGrpSize, nrOtherGroups int, firstGroup, others group, idealQE *uint64) {
	if curPkg == len(packages) {
		if len(firstGroup.pkgs) == firstGrpSize && others.weight == nrOtherGroups*firstGroup.weight &&
			makeOtherGroups(0, others.pkgs, make([]int, nrOtherGroups), firstGroup.weight) {
			*idealQE = min(*idealQE, firstGroup.QE())
		}
		return
	}

	if len(firstGroup.pkgs) < firstGrpSize {
		solve(curPkg+1, packages, firstGrpSize, nrOtherGroups, firstGroup.addPkg(packages[curPkg]), others, idealQE)
	}
	solve(curPkg+1, packages, firstGrpSize, nrOtherGroups, firstGroup, others.addPkg(packages[curPkg]), idealQE)
}

func solveFor(nrGroups int, packages []int) uint64 {
	idealQE := uint64(math.MaxUint64)
	for size := 1; size <= len(packages)-nrGroups+1; size++ {
		solve(0, packages, size, nrGroups-1, newGroup(), newGroup(), &idealQE)
		if idealQE < math.MaxUint64 {
			return idealQE
		}
	}
	return math.MaxUint64
}

func main() {
	packages := []int{}
	for {
		var pkg int
		if _, err := fmt.Scanln(&pkg); err != nil {
			break
		}
		packages = append(packages, pkg)
	}

	fmt.Println("part 1:", solveFor(3, packages))
	fmt.Println("part 2:", solveFor(4, packages))
}
