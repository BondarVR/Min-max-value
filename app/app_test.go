package app

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWorkingWithAnArray(t *testing.T) {
	//Arrange
	testTable := []struct {
		name     string
		numbers  [5]int
		expected string
	}{
		{
			name:     "First",
			numbers:  [5]int{0, 0, 0, 0, 0},
			expected: "minimum: 0, maximum: 0",
		},
		{
			name:     "Second",
			numbers:  [5]int{1, 2, 3, 4, 5},
			expected: "minimum: 10, maximum: 14",
		},
		{
			name:     "Third",
			numbers:  [5]int{100, 2250, 31548, 4, 5147},
			expected: "minimum: 7501, maximum: 39045",
		},
		{
			name:     "Fourth",
			numbers:  [5]int{0, 0, 0, 0, 5},
			expected: "minimum: 0, maximum: 5",
		},
		{
			name:     "Fifth",
			numbers:  [5]int{1, -2, 3, -4, 5},
			expected: "minimum: -2, maximum: 7",
		},
	}

	//Act
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			gotResp := WorkingWithAnArray(tt.numbers)

			assert.Equal(t, tt.expected, gotResp,
				fmt.Sprintf("Incorrect result. Test name: %s; Expected %s, got %s", tt.name, tt.expected, gotResp))
		})
	}
}
