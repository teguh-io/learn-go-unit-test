package test

import (
	"fmt"
	"sesi11/helpers"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Job(t *testing.T) {
	testCases := []struct {
		age      int
		gender   string
		expected bool
	}{
		{
			age:      16,
			gender:   "Pria",
			expected: false,
		},
		{
			age:      17,
			gender:   "Pria",
			expected: false,
		},
		{
			age:      50,
			gender:   "Pria",
			expected: true,
		},
		{
			age:      60,
			gender:   "Pria",
			expected: false,
		},
		{
			age:      16,
			gender:   "Wanita",
			expected: false,
		},
		{
			age:      21,
			gender:   "Wanita",
			expected: true,
		},
		{
			age:      50,
			gender:   "Wanita",
			expected: true,
		},
		{
			age:      60,
			gender:   "Wanita",
			expected: false,
		},
	}

	for k, testCase := range testCases {
		t.Run(fmt.Sprintf("test-%d", k+1), func(t *testing.T) {
			require.Equal(t, testCase.expected, helpers.IsMeetRequirement(testCase.age, testCase.gender))
		})
	}
}
