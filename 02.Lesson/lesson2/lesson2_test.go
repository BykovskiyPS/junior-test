package lesson2

import "testing"

type pair struct {
	first  string
	second string
}

func testIsIsomorphic(t *testing.T) {
	pairs := []pair{
		{"egg", "add"},
		{"foo", "bar"},
		{"paper", "bar"},
	}
	want := []bool{true, false, true}
	got := make([]bool, len(want))

	for i, pair := range pairs {
		got[i] = IsIsomorphic(pair.first, pair.second)
	}

	for i, v := range want {
		if v != got[i] {
			//t.Error("got %t want %t", got[i], v)
			t.Error("got", got[i], "want", v)
		}
	}
}
