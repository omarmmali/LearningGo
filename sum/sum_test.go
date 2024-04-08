package sum

import "testing"

func TestSum(t *testing.T) {
	assertEqual := func(t *testing.T, got, want int) {
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("Sum [1, 2, 3]", func(t *testing.T) {
		numbers := []int{1,2,3}
		want := 6

		got := Sum(numbers)

		assertEqual(t, got, want)
	})

	t.Run("Sum [1, 6, 3, 3]", func(t *testing.T) {
		numbers := []int{1, 6, 3, 3}
		want := 13

		got := Sum(numbers)

		assertEqual(t, got, want)
	})

	t.Run("Sum []", func(t *testing.T) {
		numbers := []int{}
		want := 0

		got := Sum(numbers)

		assertEqual(t, got, want)
	})
}