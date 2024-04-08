package dictionary

import "testing"

func TestDictionary(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}
	
	t.Run("Search for a known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStringsEqual(t, got, want)
	})

	t.Run("Search for an unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"
		assertStringsEqual(t, err.Error(), want)
	})
}

func assertStringsEqual(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
