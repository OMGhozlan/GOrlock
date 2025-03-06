package sliceoflife

import (
	"reflect"
	"testing"
)

func assertValueMessage(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSlices(t *testing.T) {
	t.Run("create a slice of length 5 with slice implementation", func(t *testing.T) {
		got := Slices()
		want := []int{1, 2, 3, 4, 5}
		assertValueMessage(t, got, want)
	})
}

func TestSumAllSlices(t *testing.T) {
	t.Run("testing sum of each of the passed slices", func(t *testing.T) {
		got := SumAllSlices([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		assertValueMessage(t, got, want)
	})
	t.Run("testing sum of each of the passed with Sum and SumAll", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		assertValueMessage(t, got, want)
	})
	t.Run("testing sum of tails each of the passed with Sum and SumAllTails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		assertValueMessage(t, got, want)
	})
	t.Run("testing sum of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}
		assertValueMessage(t, got, want)
	})
}
