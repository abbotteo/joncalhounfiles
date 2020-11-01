package testing101

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expected := 15
	actual := Sum(numbers)

	if actual != expected {
		t.Errorf("Expected the sum of %v to be %d but instead got %d!", numbers, expected, actual)
	}
}

func TestMultipleTestsForSum(t *testing.T) {
	t.Run("All positive", testSumFunc([]int{1, 2, 3, 4, 5}, 15))
	t.Run("Some Negative", testSumFunc([]int{1, -2, 3, -4, 2}, 0))
}

func testSumFunc(numbers []int, expected int) func(*testing.T) {
	return func(t *testing.T) {
		actual := Sum(numbers)
		assert.Equal(t, expected, actual, fmt.Sprintf("Values expected to match [expected: %v] - [actual: %v]", expected, actual))
	}
}
