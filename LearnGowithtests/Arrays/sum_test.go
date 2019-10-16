package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 3, 5, 6, 76, 78}
		sum := Sum(numbers)
		expected := 169

		if sum != expected {
			t.Errorf("Expected '%d' but got '%d' when given '%v'", expected, sum, numbers)
		}
	})
}
