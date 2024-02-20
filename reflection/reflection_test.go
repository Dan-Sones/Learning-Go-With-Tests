package reflection

import "testing"

func TestWalk(t *testing.T) {
	expected := "chris"

	// slice of strings that stores what strings were passed into the fn by walk
	var got []string

	// Anon struct
	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if got[0] != expected {
		t.Errorf("got %q want %q", got[0], expected)
	}

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	}

}
