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
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		word := "odin"
		definition := "awesome dog"
		dictionary.Add(word, definition)

		got, err := dictionary.Search(word)
		if err != nil {
			t.Fatalf("expected no errors but got %q", err)
		}

		assertStrings(t, got, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func assertDefinition(t testing.TB, dic Dictionary, word, definition string) {
	t.Helper()

	got, err := dic.Search(word)
	if err != nil {
		t.Fatalf("the word %s isn't in the dictionary", word)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
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
		t.Errorf("got '%+v' want '%+v'", got, want)
	}
}
