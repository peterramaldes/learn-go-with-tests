package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("knowing word", func(t *testing.T) {
		word := "test"
		got, _ := dictionary.Search(word)
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknowing word", func(t *testing.T) {
		word := "foo"
		_, err := dictionary.Search(word)
		want := ErrNotFound

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, want)
	})

}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	word := "odin"
	definition := "awesome dog"
	dictionary.Add(word, definition)

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatalf("expected no errors but got %q", err)
	}

	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
