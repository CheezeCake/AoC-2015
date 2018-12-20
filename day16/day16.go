package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type suspectList map[int]bool

type propertyValues struct {
	valueList        map[int]suspectList
	unknownValueList suspectList
}

func newPropertyValues() propertyValues {
	return propertyValues{
		valueList:        make(map[int]suspectList),
		unknownValueList: make(suspectList),
	}
}

func (pvs propertyValues) findForUnknown(id int) bool {
	_, ok := pvs.unknownValueList[id]
	return ok
}

func (pvs propertyValues) findForValue(id int, value int) bool {
	_, ok := pvs.valueList[value][id]
	return ok
}

func (pvs propertyValues) findForValuePred(id int, pred func(value int) bool) bool {
	for value, _ := range pvs.valueList {
		if pred(value) && pvs.findForValue(id, value) {
			return true
		}
	}
	return false
}

func (pvs propertyValues) Find(id int, value int) bool {
	return (pvs.findForValue(id, value) || pvs.findForUnknown(id))
}

func (pvs propertyValues) FindPred(id int, pred func(value int) bool) bool {
	return (pvs.findForValuePred(id, pred) || pvs.findForUnknown(id))
}

func initialSuspectList() suspectList {
	suspects := make(suspectList)
	for id := 1; id <= 500; id++ {
		suspects[id] = true
	}
	return suspects
}

func suspectListValue(suspects suspectList) int {
	if len(suspects) == 1 {
		for suspect, _ := range suspects {
			return suspect
		}
	}
	return -1
}

func solvePart1(mfcsam map[string]int, propertyValues map[string]propertyValues) int {
	suspects := initialSuspectList()

	for pn, pvs := range propertyValues {
		machineValue := mfcsam[pn]
		for s, _ := range suspects {
			if !pvs.Find(s, machineValue) {
				delete(suspects, s)
			}
		}
	}

	return suspectListValue(suspects)
}

func solvePart2(mfcsam map[string]int, propertyValues map[string]propertyValues) int {
	suspects := initialSuspectList()

	for pn, pvs := range propertyValues {
		machineValue := mfcsam[pn]
		for s, _ := range suspects {
			found := false
			if pn == "cats" || pn == "trees" {
				found = pvs.FindPred(s,
					func(value int) bool {
						return (value > machineValue)
					})
			} else if pn == "pomeranians" || pn == "goldfish" {
				found = pvs.FindPred(s,
					func(value int) bool {
						return (value < machineValue)
					})
			} else {
				found = pvs.Find(s, machineValue)
			}
			if !found {
				delete(suspects, s)
			}
		}
	}

	return suspectListValue(suspects)
}

func parseAunt(line string) (auntId int, properties map[string]int) {
	properties = make(map[string]int)

	space := strings.IndexRune(line, ' ')
	colon := strings.IndexRune(line, ':')
	auntId, _ = strconv.Atoi(line[space+1 : colon])

	for _, p := range strings.Split(line[colon+2:], ", ") {
		field := strings.Split(p, ": ")
		properties[field[0]], _ = strconv.Atoi(field[1])
	}

	return
}

func main() {
	propertyNames := []string{
		"children",
		"cats",
		"samoyeds",
		"pomeranians",
		"akitas",
		"vizslas",
		"goldfish",
		"trees",
		"cars",
		"perfumes",
	}
	propertyValues := make(map[string]propertyValues)
	for _, propertyName := range propertyNames {
		propertyValues[propertyName] = newPropertyValues()
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		auntId, auntProperties := parseAunt(scanner.Text())

		for propertyName, pv := range propertyValues {
			value, present := auntProperties[propertyName]
			if !present {
				pv.unknownValueList[auntId] = true
			} else {
				if pv.valueList[value] == nil {
					pv.valueList[value] = make(suspectList)
				}
				pv.valueList[value][auntId] = true
			}
		}
	}

	mfcsam := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
	fmt.Println("part 1:", solvePart1(mfcsam, propertyValues))
	fmt.Println("part 2:", solvePart2(mfcsam, propertyValues))
}
