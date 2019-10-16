package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	str := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if str != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, str)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	res := Repeat("c", 3)
	fmt.Print(res)
	// Output: ccc
}
