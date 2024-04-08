package repeat

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Repeat 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Repeat 3 times", func(t *testing.T) {
		got := Repeat("b", 3)
		want := "bbb"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
