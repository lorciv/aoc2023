package main

import "testing"

func TestHandTyp(t *testing.T) {
	tests := []struct {
		hand string
		want typ
	}{
		{"AAAAA", five},
		{"AA8AA", four},
		{"23332", full},
		{"TTT98", three},
		{"23432", twoPair},
		{"A23A4", onePair},
		{"23456", high},
	}

	for _, test := range tests {
		if got := handTyp(test.hand); got != test.want {
			t.Errorf("typ(%s) = %s, want %s", test.hand, got, test.want)
		}
	}
}
