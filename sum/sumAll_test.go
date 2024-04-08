package sum

import "testing"

func TestSumAll(t *testing.T) {
	assertEqual := func(t *testing.T, got, want []int) {
		if len(got) != len(want) {
			t.Errorf("got %v, want %v", got, want)
			return
		}

		for i := range got {
			if got[i] != want[i] {
				t.Errorf("got %v, want %v", got, want)
				return
			}
		}
	}

	t.Run("SumAll [1, 2, 3] [1, 2]", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{1, 2})
		want := []int{6, 3}

		assertEqual(t, got, want)
	})

	t.Run("SumAll [1, 2, 3]", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3})
		want := []int{6}

		assertEqual(t, got, want)
	})

	t.Run("SumAll []", func(t *testing.T) {
		got := SumAll([]int{})
		want := []int{0}

		assertEqual(t, got, want)
	})
}
func assertEqual(t *testing.T, got, want []int) {
	if len(got) != len(want) {
		t.Errorf("got %v, want %v", got, want)
		return
	}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got %v, want %v", got, want)
			return
		}
	}
}
