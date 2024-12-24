package loops

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Should return an empty string", func(t *testing.T) {
		repeated := Repeat('a', 0)
		expected := ""

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("Should return the character the specified number of times.", func(t *testing.T) {
		repeated := Repeat('a', 7)
		expected := "aaaaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

}

// BenchmarkRepeat Performance test (benchmark) in Go. It is used to measure the efficiency (such as execution time)
// of the Repeat function by running it repeatedly.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat('a', 0)
	}
}

func ExampleRepeat() {
	repeated := Repeat('b', 4)
	fmt.Println(repeated)
	// Output: bbbb
}
