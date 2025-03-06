package arrayanagrande

import (
	"testing"
)

// assertCorrectMessage checks that a value matches an expected value, and logs a
// test error if they do not match. The function is a helper for tests and is
// intended to be used with the testing.TB interface.
func assertCorrectMessage(t testing.TB, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestArrays(t *testing.T) {
	t.Run("array iteration", func(t *testing.T) {
		got := Arrays()
		want := [5]int{1, 2, 3, 4, 5}
		assertCorrectMessage(t, got, want)
	})
}

func TestSum(t *testing.T) {
	t.Run("sum of array", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertCorrectMessage(t, got, want)
	})
}
