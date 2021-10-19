package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertCorrect := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("expected %q want %q", got, want)
		}
	}

	t.Run("It returns 5 a's when given (\"a\", 5)", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertCorrect(t, repeated, expected)
	})

	t.Run("It returns 12 B's when given (\"B\", 12)", func(t *testing.T) {
		repeated := Repeat("B", 12)
		expected := "BBBBBBBBBBBB"

		assertCorrect(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("Giggity", 3)
	fmt.Println(result)
	// Output: GiggityGiggityGiggity
}
