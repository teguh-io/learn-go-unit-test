package helpers

import (
	"sort"
)

func FinalPower(fighters []int, powers []int, currentPower int) int {
	fighterPowers := map[int]int{}

	if len(fighters) != len(powers) {
		return -1
	}

	for i := 0; i < len(fighters); i++ {
		fighterPowers[fighters[i]] = powers[i]
	}

	sort.Ints(fighters)

	for _, fighter := range fighters {
		if currentPower >= fighter {
			currentPower += fighterPowers[fighter]
		} else {
			break
		}
	}

	return currentPower
}
