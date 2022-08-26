package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWorkingWithAnArray(t *testing.T) {
	//Arrange
	testTable := []struct {
		numbers  [5]int
		expected string
	}{
		{
			numbers:  [5]int{0, 0, 0, 0, 0},
			expected: "minimum: 0, maximum: 0",
		},
		{
			numbers:  [5]int{1, 2, 3, 4, 5},
			expected: "minimum: 10, maximum: 14",
		},
		{
			numbers:  [5]int{100, 2250, 31548, 4, 5147},
			expected: "minimum: 7501, maximum: 39045",
		},
		{
			numbers:  [5]int{0, 0, 0, 0, 5},
			expected: "minimum: 0, maximum: 5",
		},
		{
			numbers:  [5]int{1, -2, 3, -4, 5},
			expected: "minimum: -2, maximum: 7",
		},
	}

	//Act
	for _, value := range testTable {
		result := WorkingWithAnArray(value.numbers)

		t.Logf("Calling arr %d, result %s", value.numbers, result)

		//Assert
		assert.Equal(t, value.expected, result,
			fmt.Sprintf("Incorrect result. Expected %s, got %s", value.expected, result))
	}
}
