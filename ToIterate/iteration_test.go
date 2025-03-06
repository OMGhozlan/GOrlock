package toiterate

import (
	"testing"
)

func assertCorrectMessage(t testing.TB, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestRepeat(t *testing.T) {
	t.Run("iterate 5 times with iteration implementation", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertCorrectMessage(t, got, want)
	})
	t.Run("iterate 5 times with strings.Repeat", func(t *testing.T) {
		got := Repeat2("a", 5)
		want := "aaaaa"
		assertCorrectMessage(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
