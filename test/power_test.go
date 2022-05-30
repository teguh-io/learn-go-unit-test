package test

import (
	"fmt"
	"sesi11/helpers"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_FinalPower(t *testing.T) {
	testCases := []struct {
		fighter      []int
		power        []int
		currentPower int
		expected     int
	}{
		{
			fighter:      []int{5, 3, 9, 8},
			power:        []int{2, 2, 3, 1},
			currentPower: 3,
			expected:     7,
		},
		{
			fighter:      []int{2, 6, 3, 9},
			power:        []int{2, 2, 3, 5},
			currentPower: 2,
			expected:     14,
		},
	}

	for k, testCase := range testCases {
		t.Run(fmt.Sprintf("test-%d", k+1), func(t *testing.T) {
			require.Equal(t, testCase.expected, helpers.FinalPower(testCase.fighter, testCase.power, testCase.currentPower))
		})
	}

}
