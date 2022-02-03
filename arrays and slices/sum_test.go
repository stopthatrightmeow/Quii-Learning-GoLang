package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("A collection of 5 numbers", func(t *testing.T) {
		// Basically adds the numbers...
		// had to change from [5] to [] to change from array to slice for it to run
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	// I can test any number of numbers without specifying! :)
	t.Run("A slice of any number of numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("Got %d, wanted %d, given %v", got, want, numbers)
		}
	})
}
