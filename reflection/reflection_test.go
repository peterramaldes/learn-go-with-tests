package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	testCases := []struct {
		desc          string
		input         interface{}
		expectedCalls []string
	}{
		{
			"struct with one string field",
			struct{ Name string }{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string fields",
			struct {
				Name string
				Age  int
			}{"Chris", 3},
			[]string{"Chris"},
		},
		{
			"struct with nested string fields",
			Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		},
		{
			"pointer to things",
			&Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile{
				{33, "London"},
				{22, "Brazil"},
			},
			[]string{"London", "Brazil"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var got []string

			walk(tC.input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, tC.expectedCalls) {
				t.Errorf("got %v, want %v", got, tC.expectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
