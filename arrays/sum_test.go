package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of size 5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertEqualInts(t, got, want, numbers)
	})
	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertEqualInts(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{0, 9})
	want := []int{6, 9}

	assertEqualSlices(t, got, want)
}

func TestSumTails(t *testing.T) {
	checkSlicesSums := func(t *testing.T, got []int, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Not empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9})
		want := []int{5, 9}

		checkSlicesSums(t, got, want)
	})
	t.Run("Empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		checkSlicesSums(t, got, want)
	})
}

func TestSumStrings(t *testing.T) {
	t.Run("Single slices", func(t *testing.T) {
		got := SumStrings([]string{"a", "b", "c", "d"})
		want := "a b c d"

		assertEqualStrings(t, got, want)
	})
	t.Run("Multiple string slices", func(t *testing.T) {
		got := SumStrings([]string{"a", "b"}, []string{"c", "d"})
		want := "a b c d"

		assertEqualStrings(t, got, want)
	})
}

func assertEqualStrings(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertEqualInts(t *testing.T, got int, want int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d given %v", got, want, numbers)
	}
}

func assertEqualSlices(t *testing.T, got []int, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
