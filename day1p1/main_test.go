package main

import "testing"

func TestValue(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}
	for _, test := range tests {
		got := value(test.input)
		if got != test.want {
			t.Errorf("value(%s) = %d, want %d", test.input, got, test.want)
		}
	}
}
