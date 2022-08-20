package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []int
		expect int
		fn     func(a int, b int) int
	}{
		{
			input:  []int{1, 2, 3},
			expect: 6,
			fn: func(a int, b int) int {
				return a + b
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Reduce(testCase.input, testCase.fn, 0))
		})
	}
}
