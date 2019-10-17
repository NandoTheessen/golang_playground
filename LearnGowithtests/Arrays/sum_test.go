package arrays

import (
	"reflect"
	"testing"
)

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

func TestSummAllTails(t *testing.T) {
	got := SumAllTails([]int{}, []int{0, 9})
	want := []int{0, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%v', want '%v'", got, want)
	}
}

func BenchmarkSumAll(b *testing.B) {
	numbers1 := []int{1, 3, 5, 6}
	numbers2 := []int{4, 3, 5, 6, 6}
	for i := 0; i < b.N; i++ {
		SumAllTails(numbers1, numbers2)
	}
}
