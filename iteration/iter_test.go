package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("repeating a character n times", func(t *testing.T) {
		repeated := Repeat("a", 10, "")
		expected := "aaaaaaaaaa"
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("change all characters to uppercase", func(t *testing.T) {
		repeatedUpper := Repeat("k", 10, "U")
		expected := "KKKKKKKKKK"

		if repeatedUpper != expected {
			t.Errorf("expected %q but got %q", expected, repeatedUpper)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10, "")
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 20, "")
	fmt.Println(repeated)
	// Output: aaaaaaaaaaaaaaaaaaaa
}
